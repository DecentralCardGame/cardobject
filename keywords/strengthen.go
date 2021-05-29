package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type strengthen struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (s strengthen) Validate() error {
	return s.ValidateStruct()
}

func (s strengthen) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s strengthen) InteractionText() string {
	return "Strenghten §Target §Amount."
}

func (s strengthen) Description() string {
	return "Strengthen gives a friendly entity +X max attack and attack."
}
