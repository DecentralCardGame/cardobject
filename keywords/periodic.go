package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type periodic struct {
	Effects effects
}

func (p periodic) Resolve() cardobject.Ability {
	effects := p.Effects.Resolve()
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				TimeEventListener: &cardobject.TimeEventListener{
					TimeEvent: cardobject.TimeEvent("TICKSTART")}},
			Effects: effects}}
	return ability
}

func (p periodic) Validate() error {
	return p.ValidateStruct()
}

func (p periodic) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p periodic) InteractionText() string {
	return "At the beginning of each tick activate: §Effects."
}
