package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type battlecry struct {
	Effects effects
}

func (b battlecry) Resolve() cardobject.Ability {
	effects := b.Effects.Resolve()
	ability := cardobject.Ability{
		TriggeredAbility: &cardobject.TriggeredAbility{
			Cause: &cardobject.EventListener{
				TimeEventListener: &cardobject.TimeEventListener{
					TimeEvent: cardobject.TimeEvent(cardobject.Combat)}},
			Effects: effects}}
	return ability
}

func (b battlecry) Validate() error {
	return b.ValidateStruct()
}

func (b battlecry) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b battlecry) InteractionText() string {
	return "At the beginning of each combat activate: §Effects."
}
