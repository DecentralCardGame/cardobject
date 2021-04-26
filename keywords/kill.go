package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct{}

func (k kill) Validate() error {
	return k.ValidateStruct()
}

func (k kill) ValidateStruct() error {
	return jsonschema.ValidateStruct(k)
}

func (k kill) InteractionText() string {
	return "Kill"
}

func (k kill) Description() string {
	return "Put an opposing entity from the field in the dustpile."
}
