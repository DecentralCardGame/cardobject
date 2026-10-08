package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ascension struct{}

func (a ascension) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ascension) InteractionText() string {
	return "Ascension."
}

func (a ascension) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}

func (a ascension) Description() string {
	return "Select one friendly Entity card in your Dustpile (TARGET), move it to Void, and create one Avatar with X/X, where X is the voided card's mana cost, not its attack."
}
