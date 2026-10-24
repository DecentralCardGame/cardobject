package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountPlace struct{}

func (d discountPlace) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountPlace) InteractionText() string {
	return "Discount Place."
}

func (d discountPlace) Description() string {
	return "Reduce the mana cost of exactly one random eligible Place card in your own hand by 1."
}

func (d discountPlace) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
