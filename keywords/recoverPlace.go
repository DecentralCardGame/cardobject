package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
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
	return "Return target Place from your Dustpile to your Hand."
}

func (re recoverPlace) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
