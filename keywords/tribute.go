package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type tribute struct {
	Effects effects
}

func (t tribute) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t tribute) InteractionText() string {
	return "Tribute: §Effects."
}

func (t tribute) Description() string {
	return "Sacrifice a friendly entity to activate Effects."
}
