package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recoverEntity struct{}

func (r recoverEntity) Validate() error {
	return r.ValidateStruct()
}

func (r recoverEntity) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r recoverEntity) InteractionText() string {
	return "Recover Entity."
}

func (r recoverEntity) Description() string {
	return "Return an entity from your dustpile to your hand."
}
