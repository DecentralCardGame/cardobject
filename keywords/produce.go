package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type produce struct {
	ManaAmount cardobject.IntValue
}

func (p produce) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p produce) InteractionText() string {
	return "Produce §ManaAmount."
}

func (p produce) Description() string {
	return "Produce mana."
}

func (p produce) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY}
}
