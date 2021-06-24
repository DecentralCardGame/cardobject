package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type armor struct {
	Amount cardobject.IntValue
}

func (a armor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a armor) InteractionText() string {
	return "Armor §Amount."
}

func (a armor) Description() string {
	return "Arm gives a friendly entity +X max health and health."
}

func (a armor) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY}
}
