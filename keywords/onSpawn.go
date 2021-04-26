package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onSpawn struct {
	Effects effects
}

func (o onSpawn) Validate() error {
	return o.ValidateStruct()
}

func (o onSpawn) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onSpawn) InteractionText() string {
	return "OnSpawn: §Effects."
}

func (o onSpawn) description() string {
	return "When this spawns activate Effects."
}
