package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverPlace struct{}

func (re recoverPlace) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re recoverPlace) InteractionText() string {
	return "Recover Place."
}

func (re recoverPlace) Description() string {
	return "Return a place from your dustpile to your hand."
}
