package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type hasten struct {
	Target cardobject.CardMode
}

func (h hasten) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h hasten) InteractionText() string {
	return "Hasten §Target"
}

func (h hasten) Description() string {
	return "Allow a friendly Entity's to attack the turn it is played."
}

func (h hasten) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.ENTITYTYPE}, jsonschema.TargetMode(h.Target)
}
