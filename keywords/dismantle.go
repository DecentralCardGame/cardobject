package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dismantle struct {
	Effects effects
}

func (d dismantle) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d dismantle) InteractionText() string {
	return "Dismantle: §Effects."
}

func (d dismantle) Description() string {
	return "Sacrifice a friendly place to activate Effects."
}
