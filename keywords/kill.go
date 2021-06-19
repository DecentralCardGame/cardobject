package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct {
	Target cardobject.CardMode
}

func (k kill) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(k, r)
}

func (k kill) InteractionText() string {
	return "Kill §Target"
}

func (k kill) Description() string {
	return "Put an opposing entity from the field in the dustpile."
}
