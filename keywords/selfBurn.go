package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type selfBurn struct {
	Amount cardobject.IntValue
}

func (s selfBurn) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s selfBurn) InteractionText() string {
	return "SelfBurn §Amount."
}

func (s selfBurn) Description() string {
	return "Deal X damage to the your HQ."
}

func (s selfBurn) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY, cardobject.MYSTICISM}
}
