package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct{}

func (k kill) Resolve() cardobject.Effect {
	keyword := cardobject.Keyword(cardobject.Kill)
	zoneChange := &cardobject.EntityZoneChange{
		Zone:    cardobject.EntityZone(cardobject.Dustpile),
		Player:  cardobject.PlayerMode(cardobject.Opponent),
		Keyword: &keyword}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityZoneChange: zoneChange}}
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

func (k kill) Validate() error {
	return k.ValidateStruct()
}

func (k kill) ValidateStruct() error {
	return jsonschema.ValidateStruct(k)
}

func (k kill) InteractionText() string {
	return "Kill an entity."
}
