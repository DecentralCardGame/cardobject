package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type heal struct {
	Target cardobject.CardMode
}

func (h heal) Validate() error {
	return h.ValidateStruct()
}

func (h heal) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h heal) InteractionText() string {
	return "Heal §Target"
}

func (h heal) Description() string {
	return "Restore a friendly entities health to full life."
}
