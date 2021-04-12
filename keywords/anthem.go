package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type anthem struct {
	Tag cardobject.Tag
}

func (a anthem) Resolve() cardobject.Effect {
	simpleInt := cardobject.SimpleIntValue(1)
	intValue := cardobject.IntValue{
		SimpleIntValue: &simpleInt}
	defBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty("HEALTH"),
		IntOperator: cardobject.IntOperator("ADD"),
		IntValue:    intValue}
	attBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty("ATTACK"),
		IntOperator: cardobject.IntOperator("ADD"),
		IntValue:    intValue}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityIntManipulation: defBuff},
		cardobject.EntityManipulation{
			EntityIntManipulation: attBuff}}
	selector := cardobject.EntitySelector{
		PlayerMode: cardobject.PlayerMode("YOU"),
		CardMode:   cardobject.CardMode("ALL"),
		EntityConditions: &cardobject.EntityConditions{
			cardobject.EntityCondition{
				EntityTagCondition: &cardobject.EntityTagCondition{
					StringValue:      cardobject.SimpleStringValue(a.Tag),
					StringComparator: cardobject.StringComparator("CONTAINS")}}},
		EntityZone: cardobject.EntityZone("FIELD")}
	effect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			EntityTargetEffect: &cardobject.EntityTargetEffect{
				EntitySelector:      &selector,
				EntityManipulations: manipulations}}}
	return effect
}

func (a anthem) Validate() error {
	return a.ValidateStruct()
}

func (a anthem) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a anthem) InteractionText() string {
	return "Anthem §Tag."
}
