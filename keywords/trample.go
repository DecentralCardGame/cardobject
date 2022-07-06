package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type trample struct{}

func (t trample) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t trample) InteractionText() string {
	return "Trample."
}

func (t trample) Description() string {
	return "Deals excess damage to enemy places/HQ."
}
