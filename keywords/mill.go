package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type mill struct {
	Amount cardobject.IntValue
	Player *cardobject.PlayerMode `json:",omitempty"`
}

func (m mill) Validate() error {
	return m.ValidateStruct()
}

func (m mill) ValidateStruct() error {
	return jsonschema.ValidateStruct(m)
}

func (m mill) InteractionText() string {
	return "Mill [§Player] §Amount."
}

func (m mill) Description() string {
	return "Put cards from a player deck in the dustpile."
}
