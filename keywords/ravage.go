package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ravage struct{}

func (r ravage) ValidateType(root jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(r, root)
}

func (r ravage) InteractionText() string {
	return "Ravage."
}

func (r ravage) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}

func (r ravage) Description() string {
	return "Select one friendly Entity (TARGET) and deal its attack as damage to one random opposing Place or HQ (RANDOM)."
}
