package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type powerstone struct{}

func (p powerstone) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p powerstone) InteractionText() string {
	return "Powerstone"
}

func (p powerstone) Description() string {
	return "A 1 health place token with Pay 2: Gain 4 Wisdom. Grind This 1. and Pay 1: Harm Target 1. Grind This 1."
}
