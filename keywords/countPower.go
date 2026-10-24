package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countPower struct {
	Target cardobject.EntityMode
}

func (c countPower) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countPower) InteractionText() string {
	return "CountPower of §Target -> X."
}

func (c countPower) Description() string {
	return "Count the Attack of a friendly entity and save this number to X."
}

func (c countPower) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.NATURE}
}

func (c countPower) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.ENTITYTYPE}, jsonschema.TargetMode(c.Target)
}
