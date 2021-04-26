package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type battlecry struct {
	Effects effects
}

func (b battlecry) Validate() error {
	return b.ValidateStruct()
}

func (b battlecry) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b battlecry) InteractionText() string {
	return "Battlecry: §Effects."
}

func (b battlecry) Description() string {
	return "At the beginning of each combat activate Effects."
}
