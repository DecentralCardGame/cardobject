package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type produce struct {
	ManaAmount cardobject.SimpleIntValue
}

func (p produce) Resolve() cardobject.Effect {
	keyword := cardobject.Keyword(cardobject.Produce)
	simpleInt := cardobject.SimpleIntValue(p.ManaAmount)
	intValue := cardobject.IntValue{
		SimpleIntValue: &simpleInt}
	effect := cardobject.Effect{
		ProductionEffect: &cardobject.ProductionEffect{
			Amount:  intValue,
			Keyword: &keyword}}
	return effect
}

func (p produce) Validate() error {
	return p.ValidateStruct()
}

func (p produce) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p produce) InteractionText() string {
	return "Produce §ManaAmount."
}

func (p produce) Description() string {
	return "Produce mana."
}
