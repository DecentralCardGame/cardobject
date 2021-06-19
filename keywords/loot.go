package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type loot struct {
	Effects effects
}

func (l loot) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(l, r)
}

func (l loot) InteractionText() string {
	return "Loot: §Effects."
}

func (l loot) Description() string {
	return "When an opposing place is destroyed activate Effects."
}
