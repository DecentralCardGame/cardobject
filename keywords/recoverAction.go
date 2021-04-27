package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverAction struct{}

func (r recoverAction) Validate() error {
	return r.ValidateStruct()
}

func (r recoverAction) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r recoverAction) InteractionText() string {
	return "Recover Action."
}

func (r recoverAction) Description() string {
	return "Return an action from your dustpile to your hand."
}
