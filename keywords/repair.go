package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type repair struct {
	Amount cardobject.SimpleIntValue
}

func (r repair) Resolve() cardobject.Effect {
	xvalue := cardobject.IntValue{
		SimpleIntValue: &r.Amount}
	placeEffect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			PlaceTargetEffect: &cardobject.PlaceTargetEffect{
				PlaceSelector: &cardobject.PlaceSelector{
					PlayerMode: cardobject.PlayerMode("YOU"),
					CardMode:   cardobject.CardMode("TARGET"),
					PlaceZone:  cardobject.PlaceZone("FIELD")},
				PlaceManipulations: &cardobject.PlaceManipulations{
					cardobject.PlaceManipulation{
						PlaceIntManipulation: &cardobject.PlaceIntManipulation{
							IntProperty: cardobject.PlaceIntProperty("HEALTH"),
							IntOperator: cardobject.IntOperator("ADD"),
							IntValue:    xvalue}}}}}}
	headquarterEffect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			HeadquarterTargetEffect: &cardobject.HeadquarterTargetEffect{
				HeadquarterSelector: &cardobject.HeadquarterSelector{
					PlayerMode: cardobject.PlayerMode("YOU"),
					CardMode:   cardobject.CardMode("TARGET")},
				HeadquarterManipulations: &cardobject.HeadquarterManipulations{
					cardobject.HeadquarterManipulation{
						HeadquarterIntManipulation: &cardobject.HeadquarterIntManipulation{
							IntProperty: cardobject.HeadquarterIntProperty("HEALTH"),
							IntOperator: cardobject.IntOperator("ADD"),
							IntValue:    xvalue}}}}}}
	effect := cardobject.Effect{
		ChooseFromEffect: &cardobject.ChooseFromEffect{
			Effects: cardobject.Effects{
				placeEffect,
				headquarterEffect}}}
	return effect
}

func (r repair) Validate() error {
	return r.ValidateStruct()
}

func (r repair) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r repair) InteractionText() string {
	return "Repair §Amount."
}
