package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type selfBurn struct {
	Amount cardobject.IntValue
}

func (s selfBurn) Validate() error {
	return s.ValidateStruct()
}

func (s selfBurn) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s selfBurn) InteractionText() string {
	return "SelfBurn §Amount."
}

func (s selfBurn) Description() string {
	return "Deal X damage to the your HQ."
}
