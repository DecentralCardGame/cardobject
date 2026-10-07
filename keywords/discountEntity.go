package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountEntity struct{}

func (d discountEntity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountEntity) InteractionText() string {
	return "Discount Entity."
}

func (d discountEntity) Description() string {
	return "Reduce the mana cost of exactly one random eligible Entity card in your own hand by 1."
}

func (d discountEntity) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
