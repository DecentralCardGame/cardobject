package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type cleanse struct {
	Target cardobject.CardOppMode
}

func (c cleanse) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c cleanse) InteractionText() string {
	return "Cleanse §Target."
}

func (c cleanse) Description() string {
	return "Put selected cards from the opposing Dustpile in the Void."
}
