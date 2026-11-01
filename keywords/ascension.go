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
	return "Move a friendly Entity from the Dustpile into the Void and create an X/X Avatar with X being the voided card's mana cost."
}
