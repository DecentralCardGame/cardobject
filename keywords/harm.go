package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type harm struct {
	Target cardobject.EntityOppMode
	Amount cardobject.IntValue
}

func (h harm) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h harm) InteractionText() string {
	return "Harm §Target §Amount."
}

func (h harm) Description() string {
	return "Deal X damage to an opposing Entity."
}
