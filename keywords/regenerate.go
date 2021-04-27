package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type regenerate struct {
	Amount cardobject.SimpleIntValue
}

func (r regenerate) Validate() error {
	return r.ValidateStruct()
}

func (r regenerate) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r regenerate) InteractionText() string {
	return "Regenerate §Amount."
}

func (r regenerate) Description() string {
	return "Restore health of this entity each turn."
}
