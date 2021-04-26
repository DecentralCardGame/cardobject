package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type insight struct {
	WisdomAmount cardobject.SimpleIntValue
}

func (i insight) Validate() error {
	return i.ValidateStruct()
}

func (i insight) ValidateStruct() error {
	return jsonschema.ValidateStruct(i)
}

func (i insight) InteractionText() string {
	return "Insight §WisdomAmount."
}

func (i insight) Description() string {
	return "Gain wisdom."
}
