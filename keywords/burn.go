package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type burn struct {
	Amount cardobject.IntValue
}

func (b burn) Validate() error {
	return b.ValidateStruct()
}

func (b burn) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b burn) InteractionText() string {
	return "Burn §Amount."
}

func (b burn) Description() string {
	return "Deal X damage to the opposing HQ."
}
