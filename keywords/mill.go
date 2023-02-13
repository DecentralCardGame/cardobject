package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type mill struct {
	Amount cardobject.IntValue
	Player cardobject.PlayerMode
}

func (m mill) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(m, r)
}

func (m mill) InteractionText() string {
	return "Mill §Player §Amount."
}

func (m mill) Description() string {
	return "Put cards from a player's Deck in the Dustpile."
}

func (m mill) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.TECHNOLOGY}
}
