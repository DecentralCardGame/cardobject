package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type grow struct {
	GrowthAmount cardobject.IntValue
}

func (g grow) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(g, r)
}

func (g grow) InteractionText() string {
	return "Grow §GrowthAmount."
}

func (g grow) Description() string {
	return "Grow your HQ. 4 Growth increases its size by 1."
}

func (g grow) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.NATURE}
}
