package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type armor struct {
	Amount cardobject.IntValue
}

func (a armor) Validate() error {
	return a.ValidateStruct()
}

func (a armor) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a armor) InteractionText() string {
	return "Armor §Amount."
}

func (a armor) Description() string {
	return "Arm gives a friendly entity +X max health and health."
}
