package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onDeath struct {
	Effects effects
}

func (o onDeath) Validate() error {
	return o.ValidateStruct()
}

func (o onDeath) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onDeath) InteractionText() string {
	return "OnDeath: §Effects."
}

func (o onDeath) Description() string {
	return "When this dies activate Effects."
}
