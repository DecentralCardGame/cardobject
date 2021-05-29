package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ravage struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (r ravage) Validate() error {
	return r.ValidateStruct()
}

func (r ravage) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r ravage) InteractionText() string {
	return "Ravage §Target §Amount."
}

func (r ravage) Description() string {
	return "Deal X damage to an opposing place."
}
