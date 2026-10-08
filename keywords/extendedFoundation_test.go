package keywords

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

func TestExtendedFoundationEffects(t *testing.T) {
	amount := cardobject.SimpleIntValue(3)
	tests := []struct {
		name        string
		value       effect
		serialized  string
		interaction string
		description string
		fields      map[string]string
		classes     []jsonschema.Class
	}{
		{
			"CorpseExplosion", effect{CorpseExplosion: &corpseExplosion{Target: cardobject.All, Amount: cardobject.IntValue{SimpleIntValue: &amount}}},
			`{"CorpseExplosion":{"Target":"ALL","Amount":{"SimpleIntValue":3}}}`,
			"Corpse Explosion §Target §Amount.",
			"Select one friendly Entity card in your Dustpile (TARGET), move it to Void, and deal N damage to opposing Entities selected by Target, using Amount for N.",
			map[string]string{"Target": "EntityOppMode", "Amount": "IntValue"}, []jsonschema.Class{cardobject.MYSTICISM},
		},
		{
			"CountScrap", effect{CountScrap: &countScrap{}}, `{"CountScrap":{}}`,
			"Count Scrap -> X.", "Set X to the number of Places in your Dustpile.",
			nil, []jsonschema.Class{cardobject.TECHNOLOGY},
		},
		{
			"Rally", effect{Rally: &rally{}}, `{"Rally":{}}`,
			"Rally -> X.",
			"Set X to the number of friendly Entities on the Field that share at least one tag with this card. Count each matching Entity once, including this card if it is a friendly Entity on the Field.",
			nil, []jsonschema.Class{cardobject.CULTURE},
		},
		{
			"Autopsy", effect{Autopsy: &autopsy{}}, `{"Autopsy":{}}`,
			"Autopsy -> X.",
			"Select one friendly Entity card in your Dustpile (TARGET), move it to Void, and set X to its health. Actions and Places cannot be selected.",
			nil, []jsonschema.Class{cardobject.TECHNOLOGY},
		},
		{
			"IncrementX", effect{IncrementX: &incrementX{}}, `{"IncrementX":{}}`,
			"Increment X.", "Increase X by 1.", nil, nil,
		},
		{
			"ResetX", effect{ResetX: &resetX{}}, `{"ResetX":{}}`,
			"Reset X.", "Set X to 0.", nil, nil,
		},
		{
			"Reprogram", effect{Reprogram: &reprogram{Effect: effect{IncrementX: &incrementX{}}}},
			`{"Reprogram":{"Effect":{"IncrementX":{}}}}`,
			"Reprogram: §Effect.",
			"Choose one Action card in your own hand (TARGET). Replace one randomly selected effect on that Action with the contained Effect, rather than adding another effect.",
			map[string]string{"Effect": "effect"}, []jsonschema.Class{cardobject.TECHNOLOGY},
		},
		{
			"Requiem", effect{Requiem: &requiem{Amount: 3, Effects: effects{{ResetX: &resetX{}}}}},
			`{"Requiem":{"Amount":3,"Effects":[{"ResetX":{}}]}}`,
			"Requiem §Amount: §Effects.",
			"When resolving this effect, activate Effects only if at least N Entity cards in your own Dustpile share a tag with this card, using Amount for N. Count each matching Entity once; Actions and Places do not count.",
			map[string]string{"Amount": "BasicAmount", "Effects": "effects"}, []jsonschema.Class{cardobject.CULTURE},
		},
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.value.ValidateType(foundationRoot()); err != nil {
				t.Fatal(err)
			}
			assertFoundationJSON(t, tc.value, tc.serialized)
			var decoded effect
			if err := json.Unmarshal([]byte(tc.serialized), &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.value, decoded) {
				t.Fatalf("registration did not round-trip: %#v", decoded)
			}
			definition, err := decoded.FindImplementer()
			if err != nil {
				t.Fatal(err)
			}
			assertFoundationDefinition(t, schema, "effect", tc.name, definition, tc.interaction, tc.description, tc.fields, tc.classes)
			card := Card{Action: &action{
				CardName: "Foundation", Class: foundationRoot().ClassProvider.ClassRestriction(),
				Effects: effects{decoded}, Tags: cardobject.Tags{"GOBLIN"},
			}}
			data, err := json.Marshal(card)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := Unmarshal(data); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCriticalityDefinition(t *testing.T) {
	value := ability{Criticality: &criticality{Amount: 3, Effects: effects{{IncrementX: &incrementX{}}, {ResetX: &resetX{}}}}}
	serialized := `{"Criticality":{"Amount":3,"Effects":[{"IncrementX":{}},{"ResetX":{}}]}}`
	assertFoundationJSON(t, value, serialized)
	var decoded ability
	if err := json.Unmarshal([]byte(serialized), &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(value, decoded) {
		t.Fatalf("registration did not round-trip: %#v", decoded)
	}
	root := foundationRoot()
	if err := decoded.ValidateType(root); err != nil {
		t.Fatal(err)
	}
	if root.Keywords != 3 {
		t.Fatalf("ability and nested effects count = %d, want 3", root.Keywords)
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	assertFoundationDefinition(t, schema, "ability", "Criticality", decoded.Criticality,
		"Criticality §Amount: §Effects.",
		"At the beginning of each of your turns, if X is greater than or equal to N, activate Effects, using Amount for N. This can trigger again every turn and does not itself reset or otherwise change X.",
		map[string]string{"Amount": "BasicAmount", "Effects": "effects"}, nil)
	card := Card{Entity: &entity{
		CardName: "Criticality", Class: cardobject.Class{Culture: true},
		Abilities: abilities{decoded}, Attack: 1, Health: 1, Tags: cardobject.Tags{"HUMAN"},
	}}
	data, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Unmarshal(data); err != nil {
		t.Fatal(err)
	}
}

func TestApprovedFoundationClasses(t *testing.T) {
	amount := cardobject.SimpleIntValue(1)
	value := cardobject.IntValue{SimpleIntValue: &amount}
	tests := []struct {
		name    string
		value   jsonschema.Validateable
		classes []jsonschema.Class
	}{
		{"ToxicGrowth", effect{ToxicGrowth: &toxicGrowth{Amount: value}}, []jsonschema.Class{cardobject.NATURE}},
		{"Harden", effect{Harden: &harden{Amount: value}}, []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY}},
		{"Bloodlust", effect{Bloodlust: &bloodlust{Amount: value}}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}},
		{"Salvage", effect{Salvage: &salvage{Amount: value}}, []jsonschema.Class{cardobject.TECHNOLOGY}},
		{"MiraculousHealing", effect{MiraculousHealing: &miraculousHealing{}}, []jsonschema.Class{cardobject.NATURE}},
		{"Ambush", effect{Ambush: &ambush{Target: cardobject.Target}}, []jsonschema.Class{cardobject.NATURE, cardobject.CULTURE}},
		{"Fight", effect{Fight: &fight{Target: cardobject.Target}}, []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.CULTURE}},
		{"Ravage", effect{Ravage: &ravage{}}, []jsonschema.Class{cardobject.NATURE}},
		{"CorpseExplosion", effect{CorpseExplosion: &corpseExplosion{Target: cardobject.Target, Amount: value}}, []jsonschema.Class{cardobject.MYSTICISM}},
		{"BodyCount", effect{BodyCount: &bodyCount{}}, []jsonschema.Class{cardobject.MYSTICISM}},
		{"CountScrap", effect{CountScrap: &countScrap{}}, []jsonschema.Class{cardobject.TECHNOLOGY}},
		{"Ascension", effect{Ascension: &ascension{}}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}},
		{"Cleanse", effect{Cleanse: &cleanse{Amount: value}}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}},
		{"OnDestruction", ability{OnDestruction: &onDestruction{}}, []jsonschema.Class{cardobject.TECHNOLOGY}},
		{"OnDiscard", ability{OnDiscard: &onDiscard{}}, []jsonschema.Class{cardobject.MYSTICISM}},
		{"Rally", effect{Rally: &rally{}}, []jsonschema.Class{cardobject.CULTURE}},
		{"Requiem", effect{Requiem: &requiem{}}, []jsonschema.Class{cardobject.CULTURE}},
		{"Autopsy", effect{Autopsy: &autopsy{}}, []jsonschema.Class{cardobject.TECHNOLOGY}},
		{"Reprogram", effect{Reprogram: &reprogram{Effect: effect{ResetX: &resetX{}}}}, []jsonschema.Class{cardobject.TECHNOLOGY}},
		{"Criticality", ability{Criticality: &criticality{}}, nil},
		{"IncrementX", effect{IncrementX: &incrementX{}}, nil},
		{"ResetX", effect{ResetX: &resetX{}}, nil},
	}
	classes := map[jsonschema.Class]cardobject.Class{
		cardobject.NATURE:     {Nature: true},
		cardobject.MYSTICISM:  {Mysticism: true},
		cardobject.TECHNOLOGY: {Technology: true},
		cardobject.CULTURE:    {Culture: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for name, class := range classes {
				root := foundationRoot()
				root.ClassProvider = &action{Class: class}
				allowed := len(tc.classes) == 0 || class.Contains(tc.classes)
				if err := tc.value.ValidateType(root); (err == nil) != allowed {
					t.Errorf("class %s: allowed=%t, error=%v", name, allowed, err)
				}
			}
		})
	}
}

func TestFoundationThresholds(t *testing.T) {
	for _, n := range []cardobject.BasicAmount{-1, 0, 3, 32, 33} {
		for _, value := range []jsonschema.Validateable{
			effect{Requiem: &requiem{Amount: n, Effects: effects{{ResetX: &resetX{}}}}},
			ability{Criticality: &criticality{Amount: n, Effects: effects{{ResetX: &resetX{}}}}},
		} {
			if err := value.ValidateType(foundationRoot()); (err == nil) != (n >= 0 && n <= 32) {
				t.Errorf("threshold %d: %v", n, err)
			}
		}
	}
}

func TestCorpseExplosionAmountsAndSelectors(t *testing.T) {
	amount := cardobject.SimpleIntValue(3)
	for _, mode := range []cardobject.EntityOppMode{cardobject.Target, cardobject.Random, cardobject.All, cardobject.ChosenEntity, cardobject.This, "", "INVALID"} {
		value := effect{CorpseExplosion: &corpseExplosion{Target: mode, Amount: cardobject.IntValue{SimpleIntValue: &amount}}}
		valid := mode != cardobject.This && mode != "" && mode != "INVALID"
		if err := value.ValidateType(foundationRoot()); (err == nil) != valid {
			t.Errorf("selector %q: %v", mode, err)
		}
	}
	for _, n := range []cardobject.SimpleIntValue{-1, 0, 3, 32, 33} {
		value := effect{CorpseExplosion: &corpseExplosion{Target: cardobject.Target, Amount: cardobject.IntValue{SimpleIntValue: &n}}}
		if err := value.ValidateType(foundationRoot()); (err == nil) != (n >= 0 && n <= 32) {
			t.Errorf("amount %d: %v", n, err)
		}
	}
	for _, variable := range []cardobject.IntVariableName{cardobject.XVar, cardobject.YVar, cardobject.ZVar, "N"} {
		value := effect{CorpseExplosion: &corpseExplosion{Target: cardobject.Target, Amount: cardobject.IntValue{IntVariable: &variable}}}
		if err := value.ValidateType(foundationRoot()); (err == nil) != (variable != "N") {
			t.Errorf("variable %q: %v", variable, err)
		}
	}
	variable := cardobject.IntVariableName(cardobject.XVar)
	for _, invalid := range []cardobject.IntValue{{}, {SimpleIntValue: &amount, IntVariable: &variable}} {
		value := effect{CorpseExplosion: &corpseExplosion{Target: cardobject.Target, Amount: invalid}}
		if err := value.ValidateType(foundationRoot()); err == nil {
			t.Error("expected exactly one integer implementation")
		}
	}
}

func TestNestedFoundationValidation(t *testing.T) {
	for _, invalid := range []jsonschema.Validateable{
		effect{Reprogram: &reprogram{}},
		effect{Reprogram: &reprogram{Effect: effect{IncrementX: &incrementX{}, ResetX: &resetX{}}}},
		effect{Requiem: &requiem{Effects: effects{{}}}},
		ability{Criticality: &criticality{Effects: effects{{}}}},
		effect{Cleanse: &cleanse{}},
	} {
		if err := invalid.ValidateType(foundationRoot()); err == nil {
			t.Errorf("accepted invalid nested definition: %#v", invalid)
		}
	}
	for _, serialized := range []string{
		`{"Reprogram":{"Effect":[{"IncrementX":{}}]}}`,
		`{"Requiem":{"Amount":{"IntVariable":"X"},"Effects":[]}}`,
	} {
		var decoded effect
		if err := json.Unmarshal([]byte(serialized), &decoded); err == nil {
			t.Errorf("accepted invalid field shape: %s", serialized)
		}
	}
	var decoded ability
	if err := json.Unmarshal([]byte(`{"Criticality":{"Amount":{"IntVariable":"X"},"Effects":[]}}`), &decoded); err == nil {
		t.Error("Criticality threshold must be a fixed amount")
	}
	value := effect{IncrementX: &incrementX{}}
	for depth := 1; depth <= maxKeywords+1; depth++ {
		root := foundationRoot()
		if err := value.ValidateType(root); (err == nil) != (depth <= maxKeywords) {
			t.Errorf("nested depth %d: %v", depth, err)
		}
		if root.Keywords != depth {
			t.Errorf("nested keyword count = %d, want %d", root.Keywords, depth)
		}
		value = effect{Reprogram: &reprogram{Effect: value}}
	}
	for _, value := range []jsonschema.Validateable{
		effect{Reprogram: &reprogram{Effect: effect{Rally: &rally{}}}},
		ability{Criticality: &criticality{Effects: effects{{BodyCount: &bodyCount{}}}}},
	} {
		root := foundationRoot()
		root.ClassProvider = &action{Class: cardobject.Class{Technology: true}}
		if err := value.ValidateType(root); err == nil {
			t.Error("nested effects must retain their own class restrictions")
		}
		if err := value.ValidateType(foundationRoot()); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCountingAssignmentText(t *testing.T) {
	for _, value := range []foundationDefinition{
		count{Tag: "HUMAN"}, countDust{CardType: "ENTITY"}, countPower{Target: cardobject.Target},
		bodyCount{}, countScrap{}, rally{}, autopsy{},
	} {
		if !strings.HasSuffix(value.InteractionText(), " -> X.") {
			t.Errorf("missing X assignment: %q", value.InteractionText())
		}
	}
	for _, value := range []foundationDefinition{incrementX{}, resetX{}, sacrifice{}} {
		if strings.Contains(value.InteractionText(), "-> X") {
			t.Errorf("unexpected assignment arrow: %q", value.InteractionText())
		}
	}
}

func TestNestedFoundationSchemaSerialization(t *testing.T) {
	data, err := Schema()
	if err != nil {
		t.Fatal(err)
	}
	var schema struct {
		Definitions map[string]struct {
			Children map[string]struct {
				Ref string `json:"$ref"`
			} `json:"children"`
		} `json:"definitions"`
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		definition string
		field      string
		target     string
	}{
		{"reprogram", "Effect", "effect"},
		{"requiem", "Effects", "effects"},
		{"criticality", "Effects", "effects"},
		{"effects", "Effect", "effect"},
		{"effect", "Reprogram", "reprogram"},
		{"effect", "Requiem", "requiem"},
	} {
		if got := schema.Definitions[tc.definition].Children[tc.field].Ref; got != "#/definitions/"+tc.target {
			t.Errorf("%s.%s references %q, want %s", tc.definition, tc.field, got, tc.target)
		}
	}
	for _, removed := range []string{"Tribe", "Inscribe"} {
		if _, exists := schema.Definitions["effect"].Children[removed]; exists {
			t.Errorf("superseded keyword %s must not be registered", removed)
		}
	}
}
