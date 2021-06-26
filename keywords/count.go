package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type count struct {
	Tag cardobject.Tag
}

func (c count) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c count) InteractionText() string {
	return "Count §Tag."
}

func (c count) Description() string {
	return "Count counts up all cards on your board with a certain tag and assigns this number to X."
}
