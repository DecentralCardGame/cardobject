package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type harm struct {
	Amount cardobject.SimpleIntValue
}

func (h harm) Resolve() cardobject.Effect {
	xvalue := cardobject.IntValue{
		SimpleIntValue: &h.Amount}
	healthReduce := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty("HEALTH"),
		IntOperator: cardobject.IntOperator("SUBTRACT"),
		IntValue:    xvalue}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityIntManipulation: healthReduce}}
	selector := cardobject.EntitySelector{
		PlayerMode: cardobject.PlayerMode("OPPONENT"),
		CardMode:   cardobject.CardMode("TARGET"),
		EntityZone: cardobject.EntityZone("FIELD")}
	effect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			EntityTargetEffect: &cardobject.EntityTargetEffect{
				EntitySelector:      &selector,
				EntityManipulations: manipulations}}}
	return effect
}

func (h harm) Validate() error {
	return h.ValidateStruct()
}

func (h harm) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h harm) InteractionText() string {
	return "Harm §Amount."
}
