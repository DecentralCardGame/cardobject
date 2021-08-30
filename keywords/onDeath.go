package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onDeath struct {
	Effects effects
}

func (o onDeath) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(o, r)
}

func (o onDeath) InteractionText() string {
	return "OnDeath: §Effects."
}

func (o onDeath) Description() string {
	return "When this dies, activate Effects."
}
