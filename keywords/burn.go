package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type burn struct {
	Amount cardobject.IntValue
}

func (b burn) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b burn) InteractionText() string {
	return "Burn §Amount."
}

func (b burn) Description() string {
	return "Deal X damage to the opposing HQ."
}

func (b burn) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.CULTURE}
}
