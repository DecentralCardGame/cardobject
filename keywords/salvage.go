package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type salvage struct {
	Amount cardobject.IntValue
}

func (s salvage) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s salvage) InteractionText() string {
	return "Salvage §Amount."
}

func (s salvage) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}

func (s salvage) Description() string {
	return "Deal X damage to a friendly Place and produce X mana."
}
