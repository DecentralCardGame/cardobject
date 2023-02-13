package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type grind struct {
	Target cardobject.CardMode
	Amount cardobject.IntValue
}

func (g grind) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(g, r)
}

func (g grind) InteractionText() string {
	return "Grind §Target §Amount."
}

func (g grind) Description() string {
	return "Deal X damage to a friendly place/HQ."
}

func (g grind) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.PLACETYPE, cardobject.HEADQUARTERTYPE}, jsonschema.TargetMode(g.Target)
}
