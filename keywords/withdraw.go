package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type withdraw struct {
	Target cardobject.CardMode
}

func (w withdraw) Validate() error {
	return w.ValidateStruct()
}

func (w withdraw) ValidateStruct() error {
	return jsonschema.ValidateStruct(w)
}

func (w withdraw) InteractionText() string {
	return "Withdraw §Target."
}

func (w withdraw) Description() string {
	return "Return entities to your hand."
}
