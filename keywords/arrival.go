package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type arrival struct {
	Effects effects
}

func (a arrival) Resolve() cardobject.Ability {
	effects := a.Effects.Resolve()
	field := cardobject.Zone("FIELD")
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				ZoneChangeEventListener: &cardobject.ZoneChangeEventListener{
					Destination: &field,
					CardCondition: &cardobject.CardConditions{
						EntityConditions: &cardobject.EntityConditions{}}}},
			Effects: effects}}
	return ability
}

func (a arrival) Validate() error {
	return a.ValidateStruct()
}

func (a arrival) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a arrival) InteractionText() string {
	return "When another entity spawns activate: §Effects."
}