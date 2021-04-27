package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ravage struct {
	Amount cardobject.SimpleIntValue
}

func (r ravage) Validate() error {
	return r.ValidateStruct()
}

func (r ravage) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r ravage) InteractionText() string {
	return "Ravage §Amount."
}

func (r ravage) Description() string {
	return "Deal X damage to an opposing place."
}
