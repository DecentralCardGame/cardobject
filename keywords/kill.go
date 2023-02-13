package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct {
	Target cardobject.CardOppMode
}

func (k kill) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(k, r)
}

func (k kill) InteractionText() string {
	return "Kill §Target"
}

func (k kill) Description() string {
	return "Put an opposing Entity from the Field in the Dustpile."
}

func (k kill) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
