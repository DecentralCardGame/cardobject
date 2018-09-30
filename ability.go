package cardobject

type Ability struct {
	effect Effect
}

type ActivatedAbility struct {
	*Ability
	cost Cost
}

type TriggeredAbility struct {
	*Ability
	//event
}