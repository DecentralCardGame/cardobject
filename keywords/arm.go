package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arm struct {
	Target cardobject.CardMode
	Amount cardobject.SimpleIntValue
}

func (a arm) Validate() error {
	return a.ValidateStruct()
}

func (a arm) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a arm) InteractionText() string {
	return "Arm §Target §Amount."
}

func (a arm) Description() string {
	return "Arm gives a friendly entity +X/+X."
}
