package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type ambush struct {
	Target cardobject.CardMode
}

func (a ambush) Validate() error {
	return a.ValidateStruct()
}

func (a ambush) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a ambush) InteractionText() string {
	return "Ambush with §Target."
}

func (a ambush) Description() string {
	return "Your Entities deal damage to target opposing entity."
}
