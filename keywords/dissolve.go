package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dissolve struct {
	VoidAmount cardobject.SimpleIntValue
	Effects    effects
}

func (d dissolve) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d dissolve) InteractionText() string {
	return "Dissolve §VoidAmount: §Effects."
}

func (d dissolve) Description() string {
	return "Put cards from your Dustpile to the Void to activate Effects."
}

func (d dissolve) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}
