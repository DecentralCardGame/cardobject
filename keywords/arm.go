package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arm struct {
	Amount cardobject.SimpleIntValue
}

func (a arm) Validate() error {
	return a.ValidateStruct()
}

func (a arm) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a arm) InteractionText() string {
	return "Arm §Amount."
}

func (a arm) Description() string {
	return "Arm gives a friendly entity +X/+X."
}
