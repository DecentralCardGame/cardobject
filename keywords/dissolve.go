package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dissolve struct {
	VoidAmount cardobject.SimpleIntValue
	Effects    effects
}

func (d dissolve) Validate() error {
	return d.ValidateStruct()
}

func (d dissolve) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d dissolve) InteractionText() string {
	return "Dissolve §VoidAmount: §Effects."
}

func (d dissolve) Description() string {
	return "Put cards from your dustpile to the voids to activate Effects."
}
