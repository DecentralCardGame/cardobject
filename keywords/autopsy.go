package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type autopsy struct{}

func (a autopsy) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a autopsy) InteractionText() string {
	return "Autopsy -> X."
}

func (a autopsy) Description() string {
	return "Move a friendly Entity from your Dustpile into the Void and set X to its health."
}

func (a autopsy) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.MYSTICISM}
}
