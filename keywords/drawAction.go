package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawAction struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawAction) Validate() error {
	return d.ValidateStruct()
}

func (d drawAction) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d drawAction) InteractionText() string {
	return "Draw Action §Tag."
}

func (d drawAction) Description() string {
	return "Draw an action from your deck."
}
