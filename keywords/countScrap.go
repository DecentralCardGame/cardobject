package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type countScrap struct{}

func (c countScrap) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c countScrap) InteractionText() string {
	return "Count Scrap -> X."
}

func (c countScrap) Description() string {
	return "Set X to the number of Places in your Dustpile."
}

func (c countScrap) Classes() []jsonschema.Class {
	return []jsonschema.Class{cardobject.TECHNOLOGY}
}
