package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type recruit struct{}

func (re recruit) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(re, r)
}

func (re recruit) InteractionText() string {
	return "1/1 recruit"
}

func (re recruit) Description() string {
	return "A 1/1 entity token."
}

func (re recruit) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
