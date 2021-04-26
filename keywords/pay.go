package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type pay struct {
	ManaAmount cardobject.SimpleIntValue
	Effects    effects
}

func (p pay) Validate() error {
	return p.ValidateStruct()
}

func (p pay) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p pay) InteractionText() string {
	return "Pay §ManaAmount: §Effects."
}

func (p pay) Description() string {
	return "Pay Mana to activate Effects."
}
