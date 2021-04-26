package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onConstruction struct {
	Effects effects
}

func (o onConstruction) Resolve() cardobject.Ability {
	effects := o.Effects.Resolve()
	field := cardobject.Zone(cardobject.Field)
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				ZoneChangeEventListener: &cardobject.ZoneChangeEventListener{
					Destination: &field,
					CardCondition: &cardobject.CardConditions{
						PlaceConditions: &cardobject.PlaceConditions{}}}},
			Effects: effects}}
	return ability
}

func (o onConstruction) Validate() error {
	return o.ValidateStruct()
}

func (o onConstruction) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onConstruction) InteractionText() string {
	return "OnConstruction: §Effects."
}

func (o onConstruction) Description() string {
	return "When a friendly place spawns activate Effects."
}
