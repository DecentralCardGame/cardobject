package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type onDeath struct {
	Effects effects
}

func (o onDeath) Resolve() cardobject.Ability {
	effects := o.Effects.Resolve()
	field := cardobject.DynamicZone("FIELD")
	dustpile := cardobject.Zone("DUSTPILE")
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				ZoneChangeEventListener: &cardobject.ZoneChangeEventListener{
					Source:      &field,
					Destination: &dustpile,
					CardCondition: &cardobject.CardConditions{
						ThisConditions: &cardobject.ThisCondition{}}}},
			Effects: effects}}
	return ability
}

func (o onDeath) Validate() error {
	return o.ValidateStruct()
}

func (o onDeath) ValidateStruct() error {
	return jsonschema.ValidateStruct(o)
}

func (o onDeath) InteractionText() string {
	return "When this dies activate: §Effects."
}
