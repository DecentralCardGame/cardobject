package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type repair struct {
	Amount cardobject.SimpleIntValue
}

func (r repair) Resolve() cardobject.Effect {
	keyword := cardobject.Keyword(cardobject.Repair)
	xvalue := cardobject.IntValue{
		SimpleIntValue: &r.Amount}
	placeEffect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			PlaceTargetEffect: &cardobject.PlaceTargetEffect{
				PlaceSelector: &cardobject.PlaceSelector{
					PlayerMode: cardobject.PlayerMode(cardobject.You),
					CardMode:   cardobject.CardMode(cardobject.Target),
					PlaceZone:  cardobject.PlaceZone(cardobject.Field)},
				PlaceManipulations: &cardobject.PlaceManipulations{
					cardobject.PlaceManipulation{
						PlaceIntManipulation: &cardobject.PlaceIntManipulation{
							IntProperty: cardobject.PlaceIntProperty(cardobject.CurrentHealth),
							IntOperator: cardobject.IntOperator(cardobject.Add),
							IntValue:    xvalue,
							Keyword:     &keyword}}}}}}
	headquarterEffect := cardobject.Effect{
		TargetEffect: &cardobject.TargetEffect{
			HeadquarterTargetEffect: &cardobject.HeadquarterTargetEffect{
				HeadquarterSelector: &cardobject.HeadquarterSelector{
					PlayerMode: cardobject.PlayerMode(cardobject.You),
					CardMode:   cardobject.CardMode(cardobject.Target)},
				HeadquarterManipulations: &cardobject.HeadquarterManipulations{
					cardobject.HeadquarterManipulation{
						HeadquarterIntManipulation: &cardobject.HeadquarterIntManipulation{
							IntProperty: cardobject.HeadquarterIntProperty(cardobject.CurrentHealth),
							IntOperator: cardobject.IntOperator(cardobject.Add),
							IntValue:    xvalue,
							Keyword:     &keyword}}}}}}
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

func (r repair) Description() string {
	return "Restore lost health of a friendly place or HQ."
}
