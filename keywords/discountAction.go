package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountAction struct {
	Amount cardobject.IntValue
}

func (d discountAction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountAction) InteractionText() string {
	return "Discount Action §Amount."
}

func (d discountAction) Description() string {
	return "Reduce Mana Cost of up to X Actions in your hand by 1."
}

func (d discountAction) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
