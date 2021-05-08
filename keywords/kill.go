package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct {
	Target cardobject.CardMode
}

func (k kill) Validate() error {
	return k.ValidateStruct()
}

func (k kill) ValidateStruct() error {
	return jsonschema.ValidateStruct(k)
}

func (k kill) InteractionText() string {
	return "Kill §Target"
}

func (k kill) Description() string {
	return "Put an opposing entity from the field in the dustpile."
}
