package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bloodlust struct {
	Amount cardobject.IntValue
}

func (b bloodlust) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bloodlust) InteractionText() string {
	return "Bloodlust §Amount."
}

func (b bloodlust) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}

func (b bloodlust) Description() string {
	return "Deal X damage to a friendly Entity and X damage to a random opposing Entity."
}
