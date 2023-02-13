package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type burn struct {
	Target cardobject.CardOppMode
	Amount cardobject.IntValue
}

func (b burn) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b burn) InteractionText() string {
	return "Burn §Target §Amount."
}

func (b burn) Description() string {
	return "Deal X damage to the opposing HQ or an opposing Place."
}
