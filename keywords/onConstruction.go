package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onConstruction struct {
	Effects effects
}

func (o onConstruction) Validate() error {
	return o.ValidateStruct()
}

func (o onConstruction) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onConstruction) InteractionText() string {
	return "OnConstruction: §Effects."
}

func (o onConstruction) Description() string {
	return "When a friendly place spawns activate Effects."
}
