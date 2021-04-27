package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverPlace struct{}

func (r recoverPlace) Validate() error {
	return r.ValidateStruct()
}

func (r recoverPlace) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r recoverPlace) InteractionText() string {
	return "Recover Place."
}

func (r recoverPlace) Description() string {
	return "Return a place from your dustpile to your hand."
}
