package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type harm struct {
	Amount cardobject.SimpleIntValue
}

func (h harm) Validate() error {
	return h.ValidateStruct()
}

func (h harm) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h harm) InteractionText() string {
	return "Harm §Amount."
}

func (h harm) Description() string {
	return "Deal X damage to an opposing entity."
}
