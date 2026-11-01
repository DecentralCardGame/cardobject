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
	return "A friendly Entity deals its attack damage to an opposing Entity."
}
