package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type silence struct {
	Target cardobject.CardMode
	Player cardobject.PlayerMode
}

func (s silence) Validate() error {
	return s.ValidateStruct()
}

func (s silence) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s silence) InteractionText() string {
	return "Silence §Target §Player"
}

func (s silence) Description() string {
	return "Target entity loses its abilities."
}
