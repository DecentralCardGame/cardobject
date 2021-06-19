package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type silence struct {
	Target cardobject.CardMode
	Player cardobject.PlayerMode
}

func (s silence) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s silence) InteractionText() string {
	return "Silence §Target §Player"
}

func (s silence) Description() string {
	return "Target entity loses its abilities."
}
