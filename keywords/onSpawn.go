package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onSpawn struct {
	Effects effects
}

func (o onSpawn) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(o, r)
}

func (o onSpawn) InteractionText() string {
	return "OnSpawn: §Effects."
}

func (o onSpawn) description() string {
	return "When this spawns activate Effects."
}
