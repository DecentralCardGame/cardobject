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
	effect Effect
}

type activatedAbility struct {
	*ability
	cost Cost
	multipleUse bool
}

type triggeredAbility struct {
	*ability
	cause EventListener 
}

func (a *ability) GetSpeedModifier() SpeedModifier {
	return a.SpeedModifier
}

func (a *ability) GetEffect() Effect {
	return a.effect
}

func (aa *activatedAbility) GetCost() Cost {
	return aa.cost
}

func (aa *activatedAbility) IsMultipleUse() bool {
	return aa.multipleUse	
}

func (ta *triggeredAbility) GetEventListener() EventListener {
	return ta.cause
}