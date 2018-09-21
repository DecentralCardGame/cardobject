package cardobject

type Ability struct {
	//effect
}

type ActivatedAbility struct {
	*Ability
	cost Cost
}

type TriggeredAbility struct {
	*Ability
	//event
}