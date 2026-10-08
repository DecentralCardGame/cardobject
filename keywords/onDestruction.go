package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onDestruction struct {
	Effects effects
}

func (o onDestruction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(o, r)
}

func (o onDestruction) InteractionText() string {
	return "OnDestruction: §Effects."
}

func (o onDestruction) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}

func (o onDestruction) Description() string {
	return "Whenever another friendly Place is destroyed, including by Dismantle, activate Effects. Like Avenge, this listens for external events and never triggers on the ability bearer's own destruction."
}
