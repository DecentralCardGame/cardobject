package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bounce struct {
	Target cardobject.CardOppMode
}

func (b bounce) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bounce) InteractionText() string {
	return "Bounce §Target."
}

func (b bounce) Description() string {
	return "Return selected Entities to your opponent's hand."
}

func (b bounce) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
