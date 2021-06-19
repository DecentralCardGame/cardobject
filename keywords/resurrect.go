package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type resurrect struct{}

func (re resurrect) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re resurrect) InteractionText() string {
	return "Resurrect."
}

func (re resurrect) Description() string {
	return "Return target Entity from a dustpile to your field."
}
