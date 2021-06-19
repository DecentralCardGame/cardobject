package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type produce struct {
	ManaAmount cardobject.IntValue
}

func (p produce) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p produce) InteractionText() string {
	return "Produce §ManaAmount."
}

func (p produce) Description() string {
	return "Produce mana."
}
