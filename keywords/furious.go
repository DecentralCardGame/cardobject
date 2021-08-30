package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type furious struct {
	Effects effects
}

func (f furious) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(f, r)
}

func (f furious) InteractionText() string {
	return "Furious: §Effects."
}

func (f furious) Description() string {
	return "Whenever this takes damage, activate effects."
}
