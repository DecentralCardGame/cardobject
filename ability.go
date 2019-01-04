package cardobject

type Ability interface {
	GetSpeedModifier() SpeedModifier
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

func NewActivatedAbility(s SpeedModifier, e Effect, c Cost, u bool) ActivatedAbility {
	return &activatedAbility{&ability{s, e}, c, u}
}

func NewTriggeredAbility(s SpeedModifier, e Effect, c EventListener) TriggeredAbility {
	return &triggeredAbility{&ability{s, e}, c}
}


type ability struct {
	SpeedModifier SpeedModifier
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

func (a *ability) GetSpeedModifier() SpeedModifier {
	return a.SpeedModifier
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