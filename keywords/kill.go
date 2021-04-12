package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type kill struct{}

func (k kill) Resolve() cardobject.Effect {
	zoneChange := &cardobject.EntityZoneChange{
		Zone:   cardobject.EntityZone("DUSTPILE"),
		Player: cardobject.PlayerMode("OPPONENT")}
	manipulations := &cardobject.EntityManipulations{
		cardobject.EntityManipulation{
			EntityZoneChange: zoneChange}}
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

func (k kill) Validate() error {
	return k.ValidateStruct()
}

func (k kill) ValidateStruct() error {
	return jsonschema.ValidateStruct(k)
}

func (k kill) InteractionText() string {
	return "Kill an entity."
}
