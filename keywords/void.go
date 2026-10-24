package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type void struct {
	Target cardobject.CardOppMode
}

func (v void) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(v, r)
}

func (v void) InteractionText() string {
	return "Void §Target"
}

func (v void) Description() string {
	return "Put selected cards from your own Dustpile in the Void."
}

func (v void) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}
