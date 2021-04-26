package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dismantle struct {
	Effects effects
}

func (d dismantle) Validate() error {
	return d.ValidateStruct()
}

func (d dismantle) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d dismantle) InteractionText() string {
	return "Dismantle: §Effects."
}

func (d dismantle) Description() string {
	return "Sacrifice a friendly place to activate Effects."
}
