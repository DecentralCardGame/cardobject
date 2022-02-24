package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type fortify struct {
	Amount cardobject.IntValue
}

func (f fortify) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(f, r)
}

func (f fortify) InteractionText() string {
	return "Fortify §Amount."
}

func (f fortify) Description() string {
	return "Fortify adds X Health to friendly HQ."
}

func (f fortify) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.CULTURE}
}
