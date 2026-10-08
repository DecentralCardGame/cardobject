package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bodyCount struct{}

func (b bodyCount) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bodyCount) InteractionText() string {
	return "Body Count -> X."
}

func (b bodyCount) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}

func (b bodyCount) Description() string {
	return "Set X to the number of Entities in your Dustpile."
}
