package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type periodic struct {
	Effects effects
}

func (p periodic) Validate() error {
	return p.ValidateStruct()
}

func (p periodic) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p periodic) InteractionText() string {
	return "Periodic: §Effects."
}

func (p periodic) Description() string {
	return "At the beginning of each tick activate Effects."
}
