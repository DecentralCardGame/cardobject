package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type avenge struct {
	Effects effects
}

func (a avenge) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a avenge) InteractionText() string {
	return "Avenge: §Effects."
}

func (a avenge) Description() string {
	return "When another of your entity dies activate Effects."
}
