package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onDiscard struct {
	Effects effects
}

func (o onDiscard) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(o, r)
}

func (o onDiscard) InteractionText() string {
	return "OnDiscard: §Effects."
}

func (o onDiscard) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}

func (o onDiscard) Description() string {
	return "Whenever you discard a card, activate Effects."
}
