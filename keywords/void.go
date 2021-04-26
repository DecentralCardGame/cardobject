package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type void struct {
	Target cardobject.CardMode
}

func (v void) Validate() error {
	return v.ValidateStruct()
}

func (v void) ValidateStruct() error {
	return jsonschema.ValidateStruct(v)
}

func (v void) InteractionText() string {
	return "Void §Target"
}

func (v void) Description() string {
	return "Put opposing cards from the dustpile in the void."
}
