package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type channel struct {
	Effects effects
}

func (c channel) Validate() error {
	return c.ValidateStruct()
}

func (c channel) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c channel) InteractionText() string {
	return "Channel: §Effects."
}

func (c channel) Description() string {
	return "When you play an action activate Effects."
}
