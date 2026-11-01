package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type cleanse struct {
	Amount cardobject.IntValue
}

func (c cleanse) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c cleanse) InteractionText() string {
	return "Cleanse §Amount."
}

func (c cleanse) Description() string {
	return "N random cards from the opposing Dustpile are moved into the Void."
}

func (c cleanse) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}
