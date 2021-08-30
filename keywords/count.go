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
	return "Count all cards on your board with a selected tag and save this number to X."
}
