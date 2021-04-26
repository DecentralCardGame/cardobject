package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type harm struct {
	Amount cardobject.SimpleIntValue
}

func (h harm) Resolve() cardobject.Effect {
	keyword := cardobject.Keyword(cardobject.Harm)
	xvalue := cardobject.IntValue{
		SimpleIntValue: &h.Amount}
	healthReduce := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty(cardobject.CurrentHealth),
		IntOperator: cardobject.IntOperator(cardobject.Subtract),
		IntValue:    xvalue,
		Keyword:     &keyword}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityIntManipulation: healthReduce}}
	selector := cardobject.EntitySelector{
		PlayerMode: cardobject.PlayerMode(cardobject.Opponent),
		CardMode:   cardobject.CardMode(cardobject.Target),
		EntityZone: cardobject.EntityZone(cardobject.Field)}
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

func (h harm) Description() string {
	return "Deal X damage to an opposing entity."
}
