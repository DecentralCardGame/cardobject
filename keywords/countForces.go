package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countForces struct {
	Power cardobject.SimpleIntValue
}

func (c countForces) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countForces) InteractionText() string {
	return "Count Forces §Power."
}

func (c countForces) Description() string {
	return "Count your Entities with Attack greater or equal to selected Attack and save this number to X."
}

func (c countForces) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.NATURE}
}
