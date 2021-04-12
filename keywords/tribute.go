package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type tribute struct {
	Effects effects
}

func (t tribute) Resolve() cardobject.Ability {
	effects := t.Effects.Resolve()
	ability := cardobject.Ability{
		ActivatedAbility: &cardobject.ActivatedAbility{
			AbilityCost: &cardobject.Cost{
				SacrificeCost: &cardobject.SacrificeCost{
					Amount: cardobject.BasicAmount(1),
					Conditions: &cardobject.CardConditions{
						EntityConditions: &cardobject.EntityConditions{}}}},
			Effects: effects}}
	return ability
}

func (t tribute) Validate() error {
	return t.ValidateStruct()
}

func (t tribute) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t tribute) InteractionText() string {
	return "Sacrifice an entity to activate: §Effects."
}
