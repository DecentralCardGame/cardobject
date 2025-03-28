package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type resurrect struct {
	Amount cardobject.IntValue
}

func (re resurrect) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re resurrect) InteractionText() string {
	return "Resurrect lvl §Amount."
}

func (re resurrect) Description() string {
	return "Return target Entity from a Dustpile to your Field. The mana cost of the Entity must not be greater than the level of the Resurrect."
}

func (re resurrect) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
