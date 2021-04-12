package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arm struct {
	Amount cardobject.SimpleIntValue
}

func (a arm) Resolve() cardobject.Effect {
	xvalue := cardobject.IntValue{
		SimpleIntValue: &a.Amount}
	defBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty("HEALTH"),
		IntOperator: cardobject.IntOperator("ADD"),
		IntValue:    xvalue}
	attBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty("ATTACK"),
		IntOperator: cardobject.IntOperator("ADD"),
		IntValue:    xvalue}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityIntManipulation: defBuff},
		cardobject.EntityManipulation{
			EntityIntManipulation: attBuff}}
	selector := cardobject.EntitySelector{
		PlayerMode: cardobject.PlayerMode("YOU"),
		CardMode:   cardobject.CardMode("TARGET"),
		EntityZone: cardobject.EntityZone("FIELD")}
	effect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			EntityTargetEffect: &cardobject.EntityTargetEffect{
				EntitySelector:      &selector,
				EntityManipulations: manipulations}}}
	return effect
}

func (a arm) Validate() error {
	return a.ValidateStruct()
}

func (a arm) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a arm) InteractionText() string {
	return "Arm §Amount."
}
