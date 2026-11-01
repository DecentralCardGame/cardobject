package keywords

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type foundationDefinition interface {
	jsonschema.Validateable
	InteractionText() string
	Description() string
}

func TestFoundationEffects(t *testing.T) {
	amount := cardobject.SimpleIntValue(3)
	value := cardobject.IntValue{SimpleIntValue: &amount}
	tests := []struct {
		name        string
		effect      effect
		serialized  string
		interaction string
		description string
		fields      map[string]string
		classes     []jsonschema.Class
	}{
		{
			"ToxicGrowth", effect{ToxicGrowth: &toxicGrowth{Amount: value}},
			`{"ToxicGrowth":{"Amount":{"SimpleIntValue":3}}}`,
			"Toxic Growth §Amount.",
			"Deal X damage and give +X attack permanently to a friendly Entity.",
			map[string]string{"Amount": "IntValue"}, []jsonschema.Class{cardobject.NATURE},
		},
		{
			"Harden", effect{Harden: &harden{Amount: value}},
			`{"Harden":{"Amount":{"SimpleIntValue":3}}}`,
			"Harden §Amount.",
			"Deal X damage and give +X health permanently to a friendly Entity.",
			map[string]string{"Amount": "IntValue"}, []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY},
		},
		{
			"Bloodlust", effect{Bloodlust: &bloodlust{Amount: value}},
			`{"Bloodlust":{"Amount":{"SimpleIntValue":3}}}`,
			"Bloodlust §Amount.",
			"Deal X damage to a friendly Entity and X damage to a random opposing Entity.",
			map[string]string{"Amount": "IntValue"}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE},
		},
		{
			"Salvage", effect{Salvage: &salvage{Amount: value}},
			`{"Salvage":{"Amount":{"SimpleIntValue":3}}}`,
			"Salvage §Amount.",
			"Deal X damage to a friendly Place and produce X mana.",
			map[string]string{"Amount": "IntValue"}, []jsonschema.Class{cardobject.TECHNOLOGY},
		},
		{
			"MiraculousHealing", effect{MiraculousHealing: &miraculousHealing{}},
			`{"MiraculousHealing":{}}`,
			"Miraculous Healing.",
			"Heal a friendly Entity and increase the attack of this Entity by the amount healed.",
			nil, []jsonschema.Class{cardobject.NATURE},
		},
		{
			"Ambush", effect{Ambush: &ambush{Target: cardobject.All}},
			`{"Ambush":{"Target":"ALL"}}`,
			"Ambush §Target.",
			"A friendly Entity deals its attack damage to an opposing Entity.",
			map[string]string{"Target": "EntityOppMode"}, []jsonschema.Class{cardobject.NATURE, cardobject.CULTURE},
		},
		{
			"Fight", effect{Fight: &fight{Target: cardobject.All}},
			`{"Fight":{"Target":"ALL"}}`,
			"Fight §Target.",
			"A friendly Entity and the selected opposing Entities deal their attack damage to each other simultaneously.",
			map[string]string{"Target": "EntityOppMode"}, []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.CULTURE},
		},
		{
			"Ravage", effect{Ravage: &ravage{}},
			`{"Ravage":{}}`,
			"Ravage.",
			"Select a friendly Entity to deal its attack as damage to a random opposing Place or HQ.",
			nil, []jsonschema.Class{cardobject.NATURE},
		},
		{
			"BodyCount", effect{BodyCount: &bodyCount{}},
			`{"BodyCount":{}}`,
			"Body Count -> X.",
			"Set X to the number of Entities in your Dustpile.",
			nil, []jsonschema.Class{cardobject.NATURE, cardobject.MYSTICISM, cardobject.TECHNOLOGY, cardobject.CULTURE},
		},
		{
			"Ascension", effect{Ascension: &ascension{}},
			`{"Ascension":{}}`,
			"Ascension.",
			"Move a friendly Entity from the Dustpile into the Void and create an X/X Avatar with X being the voided card's mana cost.",
			nil, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE},
		},
		{
			"Void", effect{Void: &void{Target: cardobject.Target}},
			`{"Void":{"Target":"TARGET"}}`,
			"Void §Target",
			"Put selected cards from your own Dustpile in the Void.",
			map[string]string{"Target": "CardOppMode"}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE},
		},
		{
			"Cleanse", effect{Cleanse: &cleanse{Amount: value}},
			`{"Cleanse":{"Amount":{"SimpleIntValue":3}}}`,
			"Cleanse §Amount.",
			"N random cards from the opposing Dustpile are moved into the Void.",
			map[string]string{"Amount": "IntValue"}, []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE},
		},
		{
			"DiscountAction", effect{DiscountAction: &discountAction{}},
			`{"DiscountAction":{}}`,
			"Discount Action.",
			"Reduce the mana cost of exactly one random eligible Action card in your own hand by 1.",
			nil, []jsonschema.Class{cardobject.MYSTICISM},
		},
		{
			"DiscountEntity", effect{DiscountEntity: &discountEntity{}},
			`{"DiscountEntity":{}}`,
			"Discount Entity.",
			"Reduce the mana cost of exactly one random eligible Entity card in your own hand by 1.",
			nil, []jsonschema.Class{cardobject.CULTURE},
		},
		{
			"DiscountPlace", effect{DiscountPlace: &discountPlace{}},
			`{"DiscountPlace":{}}`,
			"Discount Place.",
			"Reduce the mana cost of exactly one random eligible Place card in your own hand by 1.",
			nil, []jsonschema.Class{cardobject.TECHNOLOGY},
		},
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.effect.ValidateType(foundationRoot()); err != nil {
				t.Fatal(err)
			}
			assertFoundationJSON(t, tc.effect, tc.serialized)
			var decoded effect
			if err := json.Unmarshal([]byte(tc.serialized), &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(decoded, tc.effect) {
				t.Fatalf("registration did not round-trip: %#v", decoded)
			}
			if err := decoded.ValidateType(foundationRoot()); err != nil {
				t.Fatal(err)
			}
			definition, err := decoded.FindImplementer()
			if err != nil {
				t.Fatal(err)
			}
			assertFoundationDefinition(t, schema, "effect", tc.name, definition, tc.interaction, tc.description, tc.fields, tc.classes)
			card := Card{Action: &action{
				CardName: "Foundation", Class: foundationRoot().ClassProvider.ClassRestriction(),
				Effects: effects{tc.effect}, Tags: cardobject.Tags{"DWARF"},
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

func TestFoundationTriggers(t *testing.T) {
	nested := effects{{BodyCount: &bodyCount{}}}
	tests := []struct {
		name        string
		ability     ability
		serialized  string
		interaction string
		description string
		classes     []jsonschema.Class
	}{
		{
			"OnDestruction", ability{OnDestruction: &onDestruction{Effects: nested}},
			`{"OnDestruction":{"Effects":[{"BodyCount":{}}]}}`,
			"OnDestruction: §Effects.",
			"Whenever another friendly Place is destroyed, activate Effects.",
			[]jsonschema.Class{cardobject.TECHNOLOGY},
		},
		{
			"OnDiscard", ability{OnDiscard: &onDiscard{Effects: nested}},
			`{"OnDiscard":{"Effects":[{"BodyCount":{}}]}}`,
			"OnDiscard: §Effects.",
			"Whenever you discard a card, activate Effects.",
			[]jsonschema.Class{cardobject.MYSTICISM},
		},
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := foundationRoot()
			if err := tc.ability.ValidateType(root); err != nil {
				t.Fatal(err)
			}
			if root.Keywords != 2 {
				t.Fatalf("trigger and nested effect count = %d, want 2", root.Keywords)
			}
			assertFoundationJSON(t, tc.ability, tc.serialized)
			var decoded ability
			if err := json.Unmarshal([]byte(tc.serialized), &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(decoded, tc.ability) {
				t.Fatalf("registration did not round-trip: %#v", decoded)
			}
			definition, err := decoded.FindImplementer()
			if err != nil {
				t.Fatal(err)
			}
			assertFoundationDefinition(t, schema, "ability", tc.name, definition, tc.interaction, tc.description, map[string]string{"Effects": "effects"}, tc.classes)
			card := Card{Entity: &entity{
				CardName: "Listener", Class: foundationRoot().ClassProvider.ClassRestriction(),
				Abilities: abilities{decoded}, Attack: 1, Health: 1, Tags: cardobject.Tags{"ELF"},
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

func foundationRoot() *CardRootValidator {
	root := allClassesTestCard()
	root.Keywords = 0
	return root
}

func assertFoundationJSON(t *testing.T, value interface{}, want string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Fatalf("JSON = %s, want %s", data, want)
	}
}

func assertFoundationDefinition(t *testing.T, schema *jsonschema.Schema, registry, name string, value jsonschema.Validateable, interaction, description string, fields map[string]string, classes []jsonschema.Class) {
	t.Helper()
	definition, ok := value.(foundationDefinition)
	if !ok {
		t.Fatalf("%s is missing definition metadata", name)
	}
	if definition.InteractionText() != interaction || definition.Description() != description {
		t.Fatalf("unexpected text for %s: %q / %q", name, definition.InteractionText(), definition.Description())
	}
	registration, ok := schema.Definitions[registry].Properties.Get(name)
	if !ok {
		t.Fatalf("%s missing from %s schema", name, registry)
	}
	ref := registration.(*jsonschema.Type).Ref
	reflected := schema.Definitions[strings.TrimPrefix(ref, "#/definitions/")]
	if reflected == nil {
		t.Fatalf("missing definition %s", ref)
	}
	if reflected.InteractionText != interaction || reflected.Description != description {
		t.Fatalf("schema metadata differs for %s", name)
	}
	if !reflect.DeepEqual(reflected.Classes, classes) {
		t.Fatalf("classes = %v, want %v", reflected.Classes, classes)
	}
	if len(reflected.Properties.Keys()) != len(fields) || len(reflected.Required) != len(fields) {
		t.Fatalf("unexpected fields or requirements: %v / %v", reflected.Properties.Keys(), reflected.Required)
	}
	for field, target := range fields {
		property, ok := reflected.Properties.Get(field)
		if !ok || property.(*jsonschema.Type).Ref != "#/definitions/"+target {
			t.Fatalf("%s must reference %s", field, target)
		}
	}
}

func TestFoundationAmounts(t *testing.T) {
	constructors := map[string]func(cardobject.IntValue) effect{
		"ToxicGrowth": func(v cardobject.IntValue) effect { return effect{ToxicGrowth: &toxicGrowth{Amount: v}} },
		"Harden":      func(v cardobject.IntValue) effect { return effect{Harden: &harden{Amount: v}} },
		"Bloodlust":   func(v cardobject.IntValue) effect { return effect{Bloodlust: &bloodlust{Amount: v}} },
		"Salvage":     func(v cardobject.IntValue) effect { return effect{Salvage: &salvage{Amount: v}} },
		"Cleanse":     func(v cardobject.IntValue) effect { return effect{Cleanse: &cleanse{Amount: v}} },
	}
	for name, build := range constructors {
		t.Run(name, func(t *testing.T) {
			for _, n := range []cardobject.SimpleIntValue{-1, 0, 3, 32, 33} {
				err := build(cardobject.IntValue{SimpleIntValue: &n}).ValidateType(foundationRoot())
				if (err == nil) != (n >= 0 && n <= 32) {
					t.Errorf("amount %d: %v", n, err)
				}
			}
			for _, variable := range []cardobject.IntVariableName{cardobject.XVar, cardobject.YVar, cardobject.ZVar, "N"} {
				value := build(cardobject.IntValue{IntVariable: &variable})
				err := value.ValidateType(foundationRoot())
				if (err == nil) != (variable != "N") {
					t.Errorf("variable %s: %v", variable, err)
				}
				assertFoundationJSON(t, value, `{"`+name+`":{"Amount":{"IntVariable":"`+string(variable)+`"}}}`)
			}
			n := cardobject.SimpleIntValue(3)
			x := cardobject.IntVariableName(cardobject.XVar)
			for _, invalid := range []cardobject.IntValue{{}, {SimpleIntValue: &n, IntVariable: &x}} {
				if err := build(invalid).ValidateType(foundationRoot()); err == nil {
					t.Error("expected exactly one integer implementation")
				}
			}
		})
	}
}

func TestFoundationSelectors(t *testing.T) {
	for _, mode := range []cardobject.EntityOppMode{cardobject.Target, cardobject.Random, cardobject.All, cardobject.ChosenEntity, cardobject.This, "", "INVALID"} {
		valid := mode != cardobject.This && mode != "" && mode != "INVALID"
		for _, value := range []effect{{Ambush: &ambush{Target: mode}}, {Fight: &fight{Target: mode}}} {
			if err := value.ValidateType(foundationRoot()); (err == nil) != valid {
				t.Errorf("Entity selector %q: %v", mode, err)
			}
		}
	}
	for _, mode := range []cardobject.CardOppMode{cardobject.Target, cardobject.Random, cardobject.All, cardobject.This, "", "INVALID"} {
		valid := mode != cardobject.This && mode != "" && mode != "INVALID"
		for _, value := range []effect{{Void: &void{Target: mode}}} {
			if err := value.ValidateType(foundationRoot()); (err == nil) != valid {
				t.Errorf("Dustpile selector %q: %v", mode, err)
			}
		}
	}
}

func TestFoundationValidationBoundaries(t *testing.T) {
	for _, invalid := range []jsonschema.Validateable{
		effect{}, effect{BodyCount: &bodyCount{}, Ravage: &ravage{}},
		ability{}, ability{OnDiscard: &onDiscard{}, OnDestruction: &onDestruction{}},
		ability{OnDiscard: &onDiscard{Effects: effects{{}}}},
		ability{OnDestruction: &onDestruction{Effects: effects{{Fight: &fight{Target: "INVALID"}}}}},
	} {
		if err := invalid.ValidateType(foundationRoot()); err == nil {
			t.Errorf("accepted invalid definition: %#v", invalid)
		}
	}
	for _, value := range []jsonschema.Validateable{
		effect{BodyCount: &bodyCount{}}, ability{OnDiscard: &onDiscard{}}, ability{OnDestruction: &onDestruction{}},
	} {
		root := foundationRoot()
		root.Keywords = maxKeywords
		if err := value.ValidateType(root); err == nil {
			t.Error("new registrations must respect the existing keyword limit")
		}
	}
	for _, value := range []effect{{DiscountAction: &discountAction{}}, {DiscountEntity: &discountEntity{}}, {DiscountPlace: &discountPlace{}}} {
		root := foundationRoot()
		root.ClassProvider = &action{Class: cardobject.Class{Nature: true}}
		if err := value.ValidateType(root); err == nil {
			t.Error("Discount must retain its class restriction")
		}
	}
}

func TestFoundationLegacyGrammar(t *testing.T) {
	for _, kind := range []cardobject.CardType{"ACTION", "ENTITY", "PLACE", "HQ"} {
		if err := (effect{CountDust: &countDust{CardType: kind}}).ValidateType(foundationRoot()); err != nil {
			t.Fatal(err)
		}
	}
	for _, value := range []effect{
		{Choose: &choose{Type: "ENTITY", Zone: "FIELD"}},
		{CountPower: &countPower{Target: cardobject.Target}},
	} {
		if err := value.ValidateType(foundationRoot()); err != nil {
			t.Fatal(err)
		}
	}
	cost := cardobject.AdditionalCost{VoidCost: &cardobject.VoidCost{Amount: 2}}
	if err := cost.ValidateType(foundationRoot()); err != nil {
		t.Fatal(err)
	}
	assertFoundationJSON(t, cost, `{"VoidCost":{"Amount":2}}`)
	if got := (dissolve{}).Description(); got != "Put cards from your Dustpile to the Void to activate Effects." {
		t.Fatalf("own Void ability cost changed: %q", got)
	}
}
