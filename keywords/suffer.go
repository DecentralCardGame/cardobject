package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type suffer struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (s suffer) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s suffer) InteractionText() string {
	return "Suffer §Target §Amount."
}

func (s suffer) Description() string {
	return "Deal X damage to a friendly Entity."
}
