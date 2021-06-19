package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onConstruction struct {
	Effects effects
}

func (o onConstruction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(o, r)
}

func (o onConstruction) InteractionText() string {
	return "OnConstruction: §Effects."
}

func (o onConstruction) Description() string {
	return "When a friendly place spawns activate Effects."
}
