package keywords

import (
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
	return "When you play an action activate Effects."
}
