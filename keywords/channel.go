package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type channel struct {
	Effects effects
}

func (c channel) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c channel) InteractionText() string {
	return "Channel: §Effects."
}

func (c channel) Description() string {
	return "Whenever you play an Action, activate Effects."
}

func (c channel) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.TECHNOLOGY}
}
