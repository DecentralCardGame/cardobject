package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type criticality struct {
	Amount  cardobject.BasicAmount
	Effects effects
}

func (c criticality) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c criticality) InteractionText() string {
	return "Criticality §Amount: §Effects."
}

func (c criticality) Description() string {
	return "At the beginning of your turn, if X is at least N, activate Effects."
}
