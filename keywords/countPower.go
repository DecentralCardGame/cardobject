package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countPower struct {
	Power cardobject.SimpleIntValue
}

func (c countPower) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countPower) InteractionText() string {
	return "Count Power §Power."
}

func (c countPower) Description() string {
	return "Count your Entities with Attack greater or equal to selected Attack."
}

func (c countPower) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.NATURE}
}
