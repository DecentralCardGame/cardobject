package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverEntity struct{}

func (re recoverEntity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re recoverEntity) InteractionText() string {
	return "Recover Entity."
}

func (re recoverEntity) Description() string {
	return "Return an entity from your dustpile to your hand."
}
