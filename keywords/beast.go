package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type beast struct{}

func (b beast) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b beast) InteractionText() string {
	return "3/3 beast"
}

func (b beast) Description() string {
	return "A 3/3 entity token."
}

func (b beast) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}
