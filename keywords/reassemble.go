package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type reassemble struct {
	Amount cardobject.IntValue
}

func (re reassemble) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re reassemble) InteractionText() string {
	return "Reassemble lvl §Amount."
}

func (re reassemble) Description() string {
	return "Return target Place from a Dustpile to your Field. The mana cost of the Place must not be greater than the level of the Reassemble."
}

func (re reassemble) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
