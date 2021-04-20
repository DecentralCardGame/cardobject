package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type anthem struct {
	Tag cardobject.Tag
}

func (a anthem) Resolve() cardobject.Effect {
	keyword := cardobject.Keyword(cardobject.Anthem)
	simpleInt := cardobject.SimpleIntValue(1)
	intValue := cardobject.IntValue{
		SimpleIntValue: &simpleInt}
	defBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty(cardobject.CurrentHealth),
		IntOperator: cardobject.IntOperator(cardobject.Add),
		IntValue:    intValue,
		Keyword:     &keyword}
	attBuff := &cardobject.EntityIntManipulation{
		IntProperty: cardobject.EntityIntProperty(cardobject.CurrentAttack),
		IntOperator: cardobject.IntOperator(cardobject.Add),
		IntValue:    intValue,
		Keyword:     &keyword}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityIntManipulation: defBuff},
		cardobject.EntityManipulation{
			EntityIntManipulation: attBuff}}
	selector := cardobject.EntitySelector{
		PlayerMode: cardobject.PlayerMode(cardobject.You),
		CardMode:   cardobject.CardMode(cardobject.All),
		EntityConditions: &cardobject.EntityConditions{
			cardobject.EntityCondition{
				EntityTagCondition: &cardobject.EntityTagCondition{
					StringValue:      cardobject.SimpleStringValue(a.Tag),
					StringComparator: cardobject.StringComparator(cardobject.Contains)}}},
		EntityZone: cardobject.EntityZone(cardobject.Field)}
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
