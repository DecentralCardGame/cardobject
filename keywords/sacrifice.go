package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type sacrifice struct {
	Target cardobject.CardMode
}

func (s sacrifice) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s sacrifice) InteractionText() string {
	return "Sacrifice §Target"
}

func (s sacrifice) Description() string {
	return "Put a friendly Entity from the Field in the Dustpile."
}
