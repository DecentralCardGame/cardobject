package cardobject

type activatedAbility struct {
	AbilityCost int8
	MultipleUse bool
	Effects     []effect
}

type triggeredAbility struct {
	Cause   eventListener
	Effects []effect
}

type ability struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}
