package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountEntity struct {
	Amount cardobject.IntValue
}

func (d discountEntity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountEntity) InteractionText() string {
	return "Discount Entity §Amount."
}

func (d discountEntity) Description() string {
	return "Reduce Mana Cost of up to X Entities in your hand by 1."
}

func (d discountEntity) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
