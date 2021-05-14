package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discard struct{}

func (d discard) Validate() error {
	return d.ValidateStruct()
}

func (d discard) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d discard) InteractionText() string {
	return "Discard."
}

func (d discard) Description() string {
	return "You have to discard a card."
}
