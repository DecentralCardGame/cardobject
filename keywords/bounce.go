package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bounce struct {
	Target cardobject.CardMode
}

func (b bounce) Validate() error {
	return b.ValidateStruct()
}

func (b bounce) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b bounce) InteractionText() string {
	return "Bounce §Target."
}

func (b bounce) Description() string {
	return "Return entities to your opponents hand."
}
