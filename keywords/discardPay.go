package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type discardPay struct {
	ManaAmount cardobject.SimpleIntValue
	Effects    effects
}

func (d discardPay) Validate() error {
	return d.ValidateStruct()
}

func (d discardPay) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d discardPay) InteractionText() string {
	return "Discard, Pay §ManaAmount: §Effects."
}

func (d discardPay) Description() string {
	return "Discard a card and pay Mana(optional) to activate Effects."
}
