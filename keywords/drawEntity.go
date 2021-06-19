package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawEntity struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawEntity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d drawEntity) InteractionText() string {
	return "Draw Entity §Tag."
}

func (d drawEntity) Description() string {
	return "Draw an entity from your deck."
}
