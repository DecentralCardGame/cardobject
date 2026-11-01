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
	return "Move a friendly Entity from your Dustpile into the Void and deal X damage to the selected opposing Entities."
}

func (c corpseExplosion) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
