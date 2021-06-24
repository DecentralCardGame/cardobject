package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type pay struct {
	ManaAmount cardobject.SimpleIntValue
	Effects    effects
}

func (p pay) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p pay) InteractionText() string {
	return "Pay §ManaAmount: §Effects."
}

func (p pay) Description() string {
	return "Pay Mana to activate Effects."
}
