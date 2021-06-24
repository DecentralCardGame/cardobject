package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverAction struct{}

func (re recoverAction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re recoverAction) InteractionText() string {
	return "Recover Action."
}

func (re recoverAction) Description() string {
	return "Return an action from your dustpile to your hand."
}

func (re recoverAction) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
