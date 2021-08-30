package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ambush struct {
	Target cardobject.CardMode
}

func (a ambush) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a ambush) InteractionText() string {
	return "Ambush with §Target."
}

func (a ambush) Description() string {
	return "The selected Entities deal damage equal to their attack to target opposing Entity."
}

func (a ambush) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}
