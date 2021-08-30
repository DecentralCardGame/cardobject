package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type train struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (t train) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t train) InteractionText() string {
	return "Train §Target §Amount."
}

func (t train) Description() string {
	return "Train gives a friendly Entity in your Hand +X/+X."
}

func (t train) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
