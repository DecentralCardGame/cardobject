package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discard struct{}

func (d discard) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d discard) InteractionText() string {
	return "Discard."
}

func (d discard) Description() string {
	return "You have to discard a card."
}
