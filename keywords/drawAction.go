package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type drawAction struct {
	Tag *cardobject.Tag `json:",omitempty"`
}

func (d drawAction) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d drawAction) InteractionText() string {
	return "Draw Action §Tag."
}

func (d drawAction) Description() string {
	return "Draw an action from your deck."
}

func (d drawAction) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM}
}
