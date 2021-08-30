package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discount struct {
	Amount cardobject.IntValue
	Type   *cardobject.CardType `json:",omitempty"`
}

func (d discount) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discount) InteractionText() string {
	return "Discount [§Type] §Amount."
}

func (d discount) Description() string {
	return "Reduce Mana Cost of all cards wit selected type in your hand."
}

func (d discount) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.TECHNOLOGY, cardobject.MYSTICISM}
}
