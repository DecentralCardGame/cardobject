package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arrival struct {
	Effects effects
}

func (a arrival) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a arrival) InteractionText() string {
	return "Arrival: §Effects."
}

func (a arrival) Description() string {
	return "When a friendly entity spawns activate Effects"
}
