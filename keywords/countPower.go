package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countPower struct {
	Target cardobject.CardMode
}

func (c countPower) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countPower) InteractionText() string {
	return "Ambush with §Target."
}

func (c countPower) Description() string {
	return "Count the Attack of a friendly entity."
}

func (c countPower) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}
