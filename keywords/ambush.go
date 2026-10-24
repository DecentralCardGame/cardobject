package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ambush struct {
	Target cardobject.EntityOppMode
}

func (a ambush) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ambush) InteractionText() string {
	return "Ambush §Target."
}

func (a ambush) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE, cardobject.CULTURE}
}

func (a ambush) Description() string {
	return "Select one friendly Entity (TARGET). Deal its attack as one-sided damage to opposing Entities selected by Target (TARGET, RANDOM, ALL, or CHOSENENTITY). There is no return damage."
}
