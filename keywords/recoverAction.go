package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverAction struct{}

func (re recoverAction) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re recoverAction) InteractionText() string {
	return "Recover Action."
}

func (re recoverAction) Description() string {
	return "Return an action from your dustpile to your hand."
}
