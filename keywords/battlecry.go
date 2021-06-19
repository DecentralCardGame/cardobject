package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type battlecry struct {
	Effects effects
}

func (b battlecry) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b battlecry) InteractionText() string {
	return "Battlecry: §Effects."
}

func (b battlecry) Description() string {
	return "At the beginning of each combat activate Effects."
}
