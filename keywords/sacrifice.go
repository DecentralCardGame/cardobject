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
	return "Select exactly one friendly Entity on the Field (TARGET), put it in the Dustpile, and set X to the sacrificed Entity's mana cost."
}
