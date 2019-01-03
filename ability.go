package cardobject

type Ability interface {
	GetEffect() Effect
}

func NewActivatedAbility(e Effect, c Cost, u bool) *ActivatedAbility {
	return &ActivatedAbility{&ability{e}, c, u}
}

func NewTriggeredAbility(e Effect, c EventListener) *TriggeredAbility {
	return &TriggeredAbility{&ability{e}, c}
}


type ability struct {
	effect Effect
}

type ActivatedAbility struct {
	*ability
	cost Cost
	multipleUse bool
}

type TriggeredAbility struct {
	*ability
	cause EventListener 
}


func (a *ability) GetEffect() Effect {
	return a.effect
}

func (aa *ActivatedAbility) GetCost() Cost {
	return aa.cost
}

func (aa *ActivatedAbility) IsMultipleUse() bool {
	return aa.multipleUse	
}

func (ta *TriggeredAbility) GetEventlistener() EventListener {
	return ta.cause
}