package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type produce struct {
	ManaAmount cardobject.IntValue
}

func (p produce) Validate() error {
	return p.ValidateStruct()
}

func (p produce) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p produce) InteractionText() string {
	return "Produce §ManaAmount."
}

func (p produce) Description() string {
	return "Produce mana."
}
