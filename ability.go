package cardobject

type activatedAbility struct {
	MultipleUse bool
	Cost []string
	Effects []effect
}

type triggeredAbility struct {
	Cause eventListener
	Effects []effect
}

type ability struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}
