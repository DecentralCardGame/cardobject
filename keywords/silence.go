package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type silence struct {
	Target cardobject.CardMode
	Player cardobject.PlayerMode
}

func (s silence) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s silence) InteractionText() string {
	return "Silence §Target §Player"
}

func (s silence) Description() string {
	return "Target Entity loses its abilities."
}

func (s silence) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
