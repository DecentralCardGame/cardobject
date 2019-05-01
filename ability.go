package cardobject

type Ability interface {
	GetEffect() Effect
}

type ActivatedAbility interface {
	Ability
	GetCost() Cost
	IsMultipleUse() bool
}

type TriggeredAbility interface {
	Ability
	GetEventListener() EventListener
}

func NewActivatedAbility(e Effect, c Cost, u bool) ActivatedAbility {
	return &activatedAbility{&ability{e}, c, u}
}

func NewTriggeredAbility(e Effect, c EventListener) TriggeredAbility {
	return &triggeredAbility{&ability{e}, c}
}


type ability struct {
	Effect Effect
}

type activatedAbility struct {
	*ability
	Cost Cost
	MultipleUse bool
}

type triggeredAbility struct {
	*ability
	Cause EventListener 
}

func (a *ability) GetEffect() Effect {
	return a.Effect
}

func (aa *activatedAbility) GetCost() Cost {
	return aa.Cost
}

func (aa *activatedAbility) IsMultipleUse() bool {
	return aa.MultipleUse	
}

func (ta *triggeredAbility) GetEventListener() EventListener {
	return ta.Cause
}