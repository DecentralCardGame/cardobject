package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type strengthen struct {
	Target cardobject.EntityMode
	Amount cardobject.IntValue
}

func (s strengthen) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s strengthen) InteractionText() string {
	return "Strengthen §Target §Amount."
}

func (s strengthen) Description() string {
	return "Strengthen gives a friendly Entity +X attack permanently."
}

func (s strengthen) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.NATURE}
}

func (s strengthen) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.ENTITYTYPE}, jsonschema.TargetMode(s.Target)
}
