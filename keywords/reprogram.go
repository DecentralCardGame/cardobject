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
	return "Choose one Action card in your own hand (TARGET). Replace one randomly selected effect on that Action with the contained Effect, rather than adding another effect."
}

func (r reprogram) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
