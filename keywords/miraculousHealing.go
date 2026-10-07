package keywords

import "github.com/DecentralCardGame/cardobject/jsonschema"

type miraculousHealing struct{}

func (m miraculousHealing) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(m, r)
}

func (m miraculousHealing) InteractionText() string {
	return "Miraculous Healing."
}

func (m miraculousHealing) Description() string {
	return "Select one friendly Entity (TARGET), heal it fully, and give that same Entity Strengthen equal to the health actually restored. A full-health Entity gains zero attack."
}
