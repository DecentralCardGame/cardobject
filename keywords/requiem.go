package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type requiem struct {
	Amount  cardobject.BasicAmount
	Effects effects
}

func (r requiem) ValidateType(root jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(r, root)
}

func (r requiem) InteractionText() string {
	return "Requiem §Amount: §Effects."
}

func (r requiem) Description() string {
	return "When resolving this effect, activate Effects only if at least N Entity cards in your own Dustpile share a tag with this card, using Amount for N. Count each matching Entity once; Actions and Places do not count."
}

func (r requiem) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
