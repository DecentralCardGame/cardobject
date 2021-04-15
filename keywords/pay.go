package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type pay struct {
	RessourceAmount cardobject.SimpleIntValue
	Effects         effects
}

func (p pay) Resolve() cardobject.Ability {
	effects := p.Effects.Resolve()
	ability := cardobject.Ability{
		ActivatedAbility: &cardobject.ActivatedAbility{
			AbilityCost: &cardobject.Cost{
				ManaCost: &cardobject.ManaCost{
					CostAmount: cardobject.BasicAmount(p.RessourceAmount)}},
			Effects: effects}}
	return ability
}

func (p pay) Validate() error {
	return p.ValidateStruct()
}

func (p pay) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p pay) InteractionText() string {
	return "Pay §RessourceAmount to activate: §Effects."
}
