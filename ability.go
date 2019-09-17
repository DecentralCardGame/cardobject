package cardobject

type ability interface {
	getAbilityAttributes() abilityAttributes
}

type abilityAttributes struct {
	Effect effect
}

type activatedAbility struct {
	abilityAttributes
	Cost []string
	MultipleUse bool
}

type triggeredAbility struct {
	abilityAttributes
	Cause eventListenerWrapper
}

type abilityWrapper struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}

func (aa abilityAttributes) getAbilityAttributes() abilityAttributes {
	return aa
}
