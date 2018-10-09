package cardobject

type Ability struct {
	effect Effect
}

type ActivatedAbility struct {
	*Ability
	cost Cost
	multipleUse bool
}

type TriggeredAbility struct {
	*Ability
	Cause EventListener 
}