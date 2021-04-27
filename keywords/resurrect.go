package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type resurrect struct{}

func (r resurrect) Validate() error {
	return r.ValidateStruct()
}

func (r resurrect) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r resurrect) InteractionText() string {
	return "Resurrect."
}

func (r resurrect) Description() string {
	return "Return target Entity from a dustpile to your field."
}
