package cardobject

type effect struct {
	ProductionEffect *productionEffect `json:",omitempty"`
	Draw             *int              `json:",omitempty"`
	TokenEffect      *tokenEffect      `json:",omitempty"`
	TargetEffect     *targetEffect     `json:",omitempty"`
}

type productionEffect struct {
	ProductionAmount int8
	ProductionType   ressources
}

type targetEffect struct {
	ActionTargetEffect *actionTargetEffect `json:",omitempty"`
	EntityTargetEffect *entityTargetEffect `json:",omitempty"`
	PlaceTargetEffect  *placeTargetEffect  `json:",omitempty"`
}

type actionTargetEffect struct {
	ActionSelector      actionSelector
	ActionManipulations []actionManipulation `json:",omitempty"`
	ZoneChange          *zoneChange          `json:",omitempty"`
}

type entityTargetEffect struct {
	EntitySelector      entitySelector
	EntityManipulations []entityManipulation `json:",omitempty"`
	ZoneChange          *zoneChange          `json:",omitempty"`
}

type placeTargetEffect struct {
	PlaceSelector      placeSelector
	PlaceManipulations []placeManipulation `json:",omitempty"`
	ZoneChange         *zoneChange         `json:",omitempty"`
}

type tokenEffect struct {
	Amount *int8 `json:",omitempty"`
	Token  token
}

type zoneChange struct {
	Zone   string
	Player string
}
