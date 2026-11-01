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

func (h harden) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.TECHNOLOGY}
}

func (h harden) Description() string {
	return "Deal X damage and give +X health permanently to a friendly Entity."
}
