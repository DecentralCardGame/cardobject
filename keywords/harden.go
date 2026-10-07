package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type harden struct {
	Amount cardobject.IntValue
}

func (h harden) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h harden) InteractionText() string {
	return "Harden §Amount."
}

func (h harden) Description() string {
	return "Select one friendly Entity (TARGET). Deal N damage to it and give that same Entity Armor N, using the same Amount for both."
}
