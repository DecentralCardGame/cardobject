package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type corpseExplosion struct {
	Target cardobject.EntityOppMode
	Amount cardobject.IntValue
}

func (c corpseExplosion) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c corpseExplosion) InteractionText() string {
	return "Corpse Explosion §Target §Amount."
}

func (c corpseExplosion) Description() string {
	return "Select one friendly Entity card in your Dustpile (TARGET), move it to Void, and deal N damage to opposing Entities selected by Target, using Amount for N."
}

func (c corpseExplosion) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
