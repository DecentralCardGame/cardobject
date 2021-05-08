package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type repair struct {
	Target cardobject.CardMode
	Amount cardobject.SimpleIntValue
}

func (r repair) Validate() error {
	return r.ValidateStruct()
}

func (r repair) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r repair) InteractionText() string {
	return "Repair §Target §Amount."
}

func (r repair) Description() string {
	return "Restore lost health of a friendly place or HQ."
}
