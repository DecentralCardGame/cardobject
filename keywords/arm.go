package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arm struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (a arm) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a arm) InteractionText() string {
	return "Arm §Target §Amount."
}

func (a arm) Description() string {
	return "Arm gives a friendly entity +X/+X."
}

func (a arm) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.MYSTICISM, cardobject.TECHNOLOGY}
}
