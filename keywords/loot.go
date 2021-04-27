package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type loot struct {
	Effects effects
}

func (l loot) Validate() error {
	return l.ValidateStruct()
}

func (l loot) ValidateStruct() error {
	return jsonschema.ValidateStruct(l)
}

func (l loot) InteractionText() string {
	return "Loot: §Effects."
}

func (l loot) Description() string {
	return "When an opposing place is destroyed activate Effects."
}
