package keywords

import (
	"encoding/json"
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

func TestSacrifice(t *testing.T) {
	sacrifice := sacrifice{}
	err := sacrifice.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	assertFoundationDefinition(t, schema, "effect", "Sacrifice", sacrifice,
		"Sacrifice.",
		"Select exactly one friendly Entity on the Field (TARGET), put it in the Dustpile, and set X to the sacrificed Entity's mana cost.",
		nil, nil)
	assertFoundationJSON(t, effect{Sacrifice: &sacrifice}, `{"Sacrifice":{}}`)
	var decoded effect
	if err := json.Unmarshal([]byte(`{"Sacrifice":{}}`), &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Sacrifice == nil {
		t.Fatal("Sacrifice registration did not round-trip")
	}
	if err := decoded.ValidateType(foundationRoot()); err != nil {
		t.Fatal(err)
	}
	assertFoundationJSON(t, decoded, `{"Sacrifice":{}}`)
	cost := cardobject.AdditionalCost{SacrificeCost: &cardobject.SacrificeCost{Amount: 2}}
	if err := cost.ValidateType(foundationRoot()); err != nil {
		t.Fatal(err)
	}
	assertFoundationJSON(t, cost, `{"SacrificeCost":{"Amount":2}}`)
}
