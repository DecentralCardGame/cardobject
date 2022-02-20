package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type tribute struct {
	Target  cardobject.CardMode
	Effects effects
}

func (t tribute) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t tribute) InteractionText() string {
	return "Tribute §Target: §Effects."
}

func (t tribute) Description() string {
	return "Sacrifice a friendly Entity to activate Effects."
}
