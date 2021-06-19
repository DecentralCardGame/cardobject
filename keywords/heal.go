package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type heal struct {
	Target cardobject.CardMode
}

func (h heal) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h heal) InteractionText() string {
	return "Heal §Target"
}

func (h heal) Description() string {
	return "Restore a friendly entities health to full life."
}
