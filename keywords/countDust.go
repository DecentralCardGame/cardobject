package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countDust struct {
	CardType cardobject.CardType `json:",omitempty"`
}

func (c countDust) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countDust) InteractionText() string {
	return "Count §CardType."
}

func (c countDust) Description() string {
	return "Count all cards of a given type in your Dustpile."
}

func (c countDust) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}