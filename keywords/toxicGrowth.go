package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type toxicGrowth struct {
	Amount cardobject.IntValue
}

func (t toxicGrowth) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t toxicGrowth) InteractionText() string {
	return "Toxic Growth §Amount."
}

func (t toxicGrowth) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}

func (t toxicGrowth) Description() string {
	return "Deal X damage and give +X attack permanently to a friendly Entity."
}
