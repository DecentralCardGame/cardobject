package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type laneswap struct {
	Target cardobject.CardMode
}

func (l laneswap) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(l, r)
}

func (l laneswap) InteractionText() string {
	return "Laneswap §Target"
}

func (l laneswap) Description() string {
	return "Chosen Entities swap their lane."
}
