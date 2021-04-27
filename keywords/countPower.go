package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countPower struct {
	Power cardobject.SimpleIntValue
}

func (c countPower) Validate() error {
	return c.ValidateStruct()
}

func (c countPower) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c countPower) InteractionText() string {
	return "Count Power §Power."
}

func (c countPower) Description() string {
	return "Count your entities with more or equal to Power."
}
