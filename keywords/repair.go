package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type repair struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (re repair) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re repair) InteractionText() string {
	return "Repair §Target §Amount."
}

func (re repair) Description() string {
	return "Restore lost health of a friendly place or HQ."
}
