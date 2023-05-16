package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type choose struct {
	Type cardobject.CardType
	Zone cardobject.DynamicZone
}

func (c choose) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c choose) InteractionText() string {
	return "Choose a §Type in §Zone."
}

func (c choose) Description() string {
	return "Choose a card which is used for future effects."
}
