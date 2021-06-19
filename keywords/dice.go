package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dice struct {
	Amount cardobject.IntValue
}

func (d dice) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d dice) InteractionText() string {
	return "Dice §Amount."
}

func (d dice) Description() string {
	return "Set X to a random number."
}
