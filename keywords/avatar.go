package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type avatar struct{}

func (a avatar) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a avatar) InteractionText() string {
	return "X/X avatar"
}

func (a avatar) Description() string {
	return "A X/X entity token."
}
