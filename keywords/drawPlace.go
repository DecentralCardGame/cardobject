package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawPlace struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawPlace) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d drawPlace) InteractionText() string {
	return "Draw Place §Tag."
}

func (d drawPlace) Description() string {
	return "Draw a place from your deck."
}
