package cardobject

type ability interface {
	getAbilityAttributes() abilityAttributes
}

type abilityAttributes struct {
	//Effect effect
}

type activatedAbility struct {
	abilityAttributes
	Cost []string
	MultipleUse bool
}

type triggeredAbility struct {
	abilityAttributes
	//Cause cause
}

func (aa abilityAttributes) getAbilityAttributes() abilityAttributes {
	return aa
}