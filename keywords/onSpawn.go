package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onSpawn struct {
	Effects effects
}

func (o onSpawn) Resolve() cardobject.Ability {
	effects := o.Effects.Resolve()
	field := cardobject.Zone("FIELD")
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				ZoneChangeEventListener: &cardobject.ZoneChangeEventListener{
					Destination: &field,
					CardCondition: &cardobject.CardConditions{
						ThisConditions: &cardobject.ThisCondition{}}}},
			Effects: effects}}
	return ability
}

func (o onSpawn) Validate() error {
	return o.ValidateStruct()
}

func (o onSpawn) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onSpawn) InteractionText() string {
	return "When this spawns activate: §Effects."
}
