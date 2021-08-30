package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type loot struct {
	Effects effects
}

func (l loot) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(l, r)
}

func (l loot) InteractionText() string {
	return "Loot: §Effects."
}

func (l loot) Description() string {
	return "Whenever an opposing Place is destroyed, activate Effects."
}

func (l loot) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.NATURE}
}
