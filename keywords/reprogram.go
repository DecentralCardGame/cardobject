package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type reprogram struct {
	Effect effect
}

func (r reprogram) ValidateType(root jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(r, root)
}

func (r reprogram) InteractionText() string {
	return "Reprogram: §Effect."
}

func (r reprogram) Description() string {
	return "Replace a random effect on an Action card in your hand with a specified Effect."
}

func (r reprogram) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
