package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discountPlace struct {
	Amount cardobject.IntValue
}

func (d discountPlace) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discountPlace) InteractionText() string {
	return "Discount Place §Amount."
}

func (d discountPlace) Description() string {
	return "Reduce Mana Cost of up to X Places in your hand by 1."
}

func (d discountPlace) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
