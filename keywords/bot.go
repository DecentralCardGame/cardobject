package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bot struct{}

func (b bot) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bot) InteractionText() string {
	return "2/2 Bot"
}

func (b bot) Description() string {
	return "A 2/2 entity token."
}

func (b bot) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
