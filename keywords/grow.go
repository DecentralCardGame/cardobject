package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type grow struct {
	GrowthAmount cardobject.SimpleIntValue
}

func (g grow) Validate() error {
	return g.ValidateStruct()
}

func (g grow) ValidateStruct() error {
	return jsonschema.ValidateStruct(g)
}

func (g grow) InteractionText() string {
	return "Grow §GrowthAmount."
}

func (g grow) Description() string {
	return "Grow your HQ."
}
