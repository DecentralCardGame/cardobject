package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type fight struct {
	Target cardobject.EntityOppMode
}

func (f fight) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(f, r)
}

func (f fight) InteractionText() string {
	return "Fight §Target."
}

func (f fight) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY, cardobject.CULTURE}
}

func (f fight) Description() string {
	return "A friendly Entity and the selected opposing Entities deal their attack damage to each other simultaneously."
}
