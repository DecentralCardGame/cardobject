package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawEntity struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawEntity) Validate() error {
	return d.ValidateStruct()
}

func (d drawEntity) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d drawEntity) InteractionText() string {
	return "Draw Entity §Tag."
}

func (d drawEntity) Description() string {
	return "Draw an entity from your deck."
}
