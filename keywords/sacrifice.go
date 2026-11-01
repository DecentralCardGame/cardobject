package keywords

import "github.com/DecentralCardGame/cardobject/jsonschema"

type sacrifice struct{}

func (s sacrifice) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s sacrifice) InteractionText() string {
	return "Sacrifice."
}

func (s sacrifice) Description() string {
	return "Select a friendly Entity, it dies. Set X to its mana cost."
}
