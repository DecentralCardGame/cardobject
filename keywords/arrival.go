package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
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
	return "When a friendly Entity is spawned, activate Effects"
}

func (a arrival) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.NATURE}
}
