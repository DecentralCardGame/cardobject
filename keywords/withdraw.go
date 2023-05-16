package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type withdraw struct {
	Target cardobject.EntityMode
}

func (w withdraw) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(w, r)
}

func (w withdraw) InteractionText() string {
	return "Withdraw §Target."
}

func (w withdraw) Description() string {
	return "Return Entities from the Field to your Hand."
}

func (w withdraw) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.MYSTICISM, cardobject.CULTURE}
}

func (w withdraw) Targets() ([]string, jsonschema.TargetMode) {
	return []string{cardobject.ENTITYTYPE}, jsonschema.TargetMode(w.Target)
}
