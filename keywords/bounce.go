package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bounce struct {
	Target cardobject.CardMode
}

func (b bounce) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bounce) InteractionText() string {
	return "Bounce §Target."
}

func (b bounce) Description() string {
	return "Return entities to your opponents hand."
}
