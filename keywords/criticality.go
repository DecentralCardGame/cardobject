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
	return "At the beginning of each of your turns, if X is greater than or equal to N, activate Effects, using Amount for N. This can trigger again every turn and does not itself reset or otherwise change X."
}
