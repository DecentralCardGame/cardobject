package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discount struct {
	Amount cardobject.IntValue
	Type   *cardobject.CardType `json:",omitempty"`
}

func (d discount) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discount) InteractionText() string {
	return "Discount [§Type] §Amount."
}

func (d discount) Description() string {
	return "Reduce manacost of all cards of a type in your hand."
}
