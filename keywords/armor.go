package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type armor struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (a armor) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a armor) InteractionText() string {
	return "Armor §Target §Amount."
}

func (a armor) Description() string {
	return "Arm gives a friendly Entity +X Health permanently."
}

func (a armor) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY}
}
