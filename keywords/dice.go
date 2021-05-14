package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dice struct {
	Amount cardobject.SimpleIntValue
}

func (d dice) Validate() error {
	return d.ValidateStruct()
}

func (d dice) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d dice) InteractionText() string {
	return "Dice §Amount."
}

func (d dice) Description() string {
	return "Set X to a random number."
}
