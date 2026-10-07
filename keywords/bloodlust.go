package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type bloodlust struct {
	Amount cardobject.IntValue
}

func (b bloodlust) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b bloodlust) InteractionText() string {
	return "Bloodlust §Amount."
}

func (b bloodlust) Description() string {
	return "Deal N damage to one selected friendly Entity (TARGET) and N damage to one random opposing Entity (RANDOM), using the same Amount for both."
}
