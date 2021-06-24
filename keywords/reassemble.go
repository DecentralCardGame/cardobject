package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type reassemble struct{}

func (re reassemble) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re reassemble) InteractionText() string {
	return "Reassemle."
}

func (re reassemble) Description() string {
	return "Return target Place from a dustpile to your field."
}

func (re reassemble) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
