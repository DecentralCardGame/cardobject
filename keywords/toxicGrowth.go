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

func (t toxicGrowth) Description() string {
	return "Select one friendly Entity (TARGET). Deal N damage to it and give that same Entity Strengthen N, using the same Amount for both."
}
