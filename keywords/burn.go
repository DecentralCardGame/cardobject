package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type burn struct {
	Amount cardobject.IntValue
}

func (b burn) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b burn) InteractionText() string {
	return "Burn §Amount."
}

func (b burn) Description() string {
	return "Deal X damage to the opposing HQ."
}
