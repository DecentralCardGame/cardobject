package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arrival struct {
	Effects effects
}

func (a arrival) Validate() error {
	return a.ValidateStruct()
}

func (a arrival) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a arrival) InteractionText() string {
	return "Arrival: §Effects."
}

func (a arrival) Description() string {
	return "When a friendly entity spawns activate Effects"
}
