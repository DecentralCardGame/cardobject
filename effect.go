package cardobject

type effect struct {
	Production *ressources `json:",omitempty"`
	Draw *int `json:",omitempty"`
	TokenEffect *tokenEffect `json:",omitempty"`
	TargetEffect *targetEffect `json:",omitempty"`
}

type targetEffect struct {
	ActionTargetEffect *actionTargetEffect `json:",omitempty"`
	EntityTargetEffect *entityTargetEffect `json:",omitempty"`
	LocationTargetEffect *locationTargetEffect `json:",omitempty"`
}

type actionTargetEffect struct {
	ActionSelector actionSelector
	ActionManipulations []actionManipulation `json:",omitempty"`
	ZoneChange *zoneChange `json:",omitempty"`
}

type entityTargetEffect struct {
	EntitySelector entitySelector
	EntityManipulations []entityManipulation `json:",omitempty"`
	ZoneChange *zoneChange `json:",omitempty"`
}

type locationTargetEffect struct {
	LocationSelector locationSelector
	LocationManipulations []locationManipulation `json:",omitempty"`
	ZoneChange *zoneChange `json:",omitempty"`
}

type tokenEffect struct {
	Amount *int8 `json:",omitempty"`
	Token token
}

type zoneChange struct {
	Zone string
	Player string
}
