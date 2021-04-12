package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type produce struct {
	RessourceAmount cardobject.SimpleIntValue
}

func (p produce) Resolve() cardobject.Effect {
	simpleInt := cardobject.SimpleIntValue(p.RessourceAmount)
	intValue := cardobject.IntValue{
		SimpleIntValue: &simpleInt}
	effect := cardobject.Effect{
		ProductionEffect: &cardobject.ProductionEffect{
			Amount: intValue}}
	return effect
}

func (p produce) Validate() error {
	return p.ValidateStruct()
}

func (p produce) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p produce) InteractionText() string {
	return "Produce §RessourceAmount ressources."
}
