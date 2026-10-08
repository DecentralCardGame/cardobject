package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

func TestSacrifice(t *testing.T) {
	sacrifice := sacrifice{"TARGET"}
	err := sacrifice.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
	schema, err := jsonschema.Reflect(&Card{})
	if err != nil {
		t.Fatal(err)
	}
	assertFoundationDefinition(t, schema, "effect", "Sacrifice", sacrifice,
		"Sacrifice §Target",
		"Put a friendly Entity from the Field in the Dustpile and set X to the sacrificed Entity's mana cost.",
		map[string]string{"Target": "EntityMode"}, nil)
	assertFoundationJSON(t, effect{Sacrifice: &sacrifice}, `{"Sacrifice":{"Target":"TARGET"}}`)
	cost := cardobject.AdditionalCost{SacrificeCost: &cardobject.SacrificeCost{Amount: 2}}
	if err := cost.ValidateType(foundationRoot()); err != nil {
		t.Fatal(err)
	}
	assertFoundationJSON(t, cost, `{"SacrificeCost":{"Amount":2}}`)
}
