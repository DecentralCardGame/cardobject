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
	return "If at least N Entities in your Dustpile share a tag with this card, activate Effects."
}

func (r requiem) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.CULTURE}
}
