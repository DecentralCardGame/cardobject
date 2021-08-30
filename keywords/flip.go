package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type flip struct {
	HeadEffect effect
	TailEffect effect
}

func (f flip) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(f, r)
}

func (f flip) InteractionText() string {
	return "Flip §HeadEffect §TailEffect."
}

func (f flip) Description() string {
	return "Randomly choose one of two effects."
}
