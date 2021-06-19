package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type anthem struct {
	Tag cardobject.Tag
}

func (a anthem) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a anthem) InteractionText() string {
	return "Anthem §Tag."
}

func (a anthem) Description() string {
	return "Anthem gives all friendly entities with a certain tag +1/+1."
}
