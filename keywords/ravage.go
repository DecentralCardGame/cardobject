package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ravage struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (ra ravage) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(ra, r)
}

func (ra ravage) InteractionText() string {
	return "Ravage §Target §Amount."
}

func (ra ravage) Description() string {
	return "Deal X damage to an opposing place."
}
