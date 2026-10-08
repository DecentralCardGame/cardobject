package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type fight struct {
	Target cardobject.EntityOppMode
}

func (f fight) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(f, r)
}

func (f fight) InteractionText() string {
	return "Fight §Target."
}

func (f fight) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.CULTURE}
}

func (f fight) Description() string {
	return "Select one friendly Entity (TARGET) and opposing Entities via Target (TARGET, RANDOM, ALL, or CHOSENENTITY). Both sides deal damage equal to their respective attack values, sampled before damage and applied simultaneously. For ALL, every opposing Entity receives the friendly Entity's attack, and the friendly Entity receives every opposing Entity's attack. Lethally hit participants still deal damage."
}
