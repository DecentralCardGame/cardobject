package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type reassemble struct{}

func (r reassemble) Validate() error {
	return r.ValidateStruct()
}

func (r reassemble) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r reassemble) InteractionText() string {
	return "Reassemle."
}

func (r reassemble) Description() string {
	return "Return target Place from a dustpile to your field."
}
