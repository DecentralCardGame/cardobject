package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type insight struct {
	WisdomAmount cardobject.IntValue
}

func (i insight) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(i, r)
}

func (i insight) InteractionText() string {
	return "Insight §WisdomAmount."
}

func (i insight) Description() string {
	return "Gain Wisdom. 4 Wisdom draws you a card."
}
