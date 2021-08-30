package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type tag struct {
	Tag    cardobject.Tag
	Player *cardobject.PlayerMode `json:",omitempty"`
}

func (t tag) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (tag) InteractionText() string {
	return "Tag [§Player] §Tag."
}

func (t tag) Description() string {
	return "Tag adds a selected Tag to target card."
}
