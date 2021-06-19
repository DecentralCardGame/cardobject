package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type mill struct {
	Amount cardobject.IntValue
	Player *cardobject.PlayerMode `json:",omitempty"`
}

func (m mill) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(m, r)
}

func (m mill) InteractionText() string {
	return "Mill [§Player] §Amount."
}

func (m mill) Description() string {
	return "Put cards from a player deck in the dustpile."
}
