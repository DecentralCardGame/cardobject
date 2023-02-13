package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type enrage struct {
	Target cardobject.CardMode
}

func (e enrage) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e enrage) InteractionText() string {
	return "Enrage §Target"
}

func (e enrage) Description() string {
	return "Enraged Entity deals excess combat damage to places."
}
