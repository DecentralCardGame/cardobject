package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type withdraw struct {
	Target cardobject.CardMode
}

func (w withdraw) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(w, r)
}

func (w withdraw) InteractionText() string {
	return "Withdraw §Target."
}

func (w withdraw) Description() string {
	return "Return entities to your hand."
}
