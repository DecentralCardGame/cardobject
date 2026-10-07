package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountAction struct{}

func (d discountAction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountAction) InteractionText() string {
	return "Discount Action."
}

func (d discountAction) Description() string {
	return "Reduce the mana cost of exactly one random eligible Action card in your own hand by 1."
}

func (d discountAction) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
