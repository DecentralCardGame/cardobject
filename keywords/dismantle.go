package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type dismantle struct {
	Target  cardobject.CardMode
	Effects effects
}

func (d dismantle) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d dismantle) InteractionText() string {
	return "Dismantle §Target: §Effects."
}

func (d dismantle) Description() string {
	return "Sacrifice a friendly Place to activate Effects."
}

func (d dismantle) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE, cardobject.TECHNOLOGY}
}
