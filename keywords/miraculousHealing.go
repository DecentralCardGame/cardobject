package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type miraculousHealing struct{}

func (m miraculousHealing) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(m, r)
}

func (m miraculousHealing) InteractionText() string {
	return "Miraculous Healing."
}

func (m miraculousHealing) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}

func (m miraculousHealing) Description() string {
	return "Heal a friendly Entity and increase the attack of this Entity by the amount healed."
}
