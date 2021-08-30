package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawPlace struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawPlace) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d drawPlace) InteractionText() string {
	return "Draw Place §Tag."
}

func (d drawPlace) Description() string {
	return "Draw a Place from your deck."
}

func (d drawPlace) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
