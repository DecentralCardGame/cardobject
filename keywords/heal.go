package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type heal struct {
	Target cardobject.EntityMode
}

func (h heal) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h heal) InteractionText() string {
	return "Heal §Target"
}

func (h heal) Description() string {
	return "Restore a friendly Entity's health to full life."
}

func (h heal) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}

func (h heal) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.ENTITYTYPE}, jsonschema.TargetMode(h.Target)
}
