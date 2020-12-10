package cardobject

type actionSelector struct {
	PlayerMode       string
	PlayerCondition  *playerConditionInterface `json:",omitempty"`
	CardMode         string
	ActionConditions *actionConditions `json:",omitempty"`
	Zone             string
	ActionExtractors *actionExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

type entitySelector struct {
	PlayerMode       string
	PlayerCondition  *playerConditionInterface `json:",omitempty"`
	CardMode         string
	EntityConditions *entityConditions `json:",omitempty"`
	Zone             string
	EntityExtractors *entityExtractors `json:",omitempty"`
	AmountExtractor  *intExtractor     `json:",omitempty"`
}

type placeSelector struct {
	PlayerMode      string
	PlayerCondition *playerConditionInterface `json:",omitempty"`
	CardMode        string
	PlaceConditions *placeConditions `json:",omitempty"`
	Zone            string
	PlaceExtractors *placeExtractors `json:",omitempty"`
	AmountExtractor *intExtractor    `json:",omitempty"`
}

func (s *actionSelector) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlayerMode(s.PlayerMode))
	if s.PlayerCondition != nil {
		errorRange = append(errorRange, s.PlayerCondition.validate())
	}
	errorRange = append(errorRange, validateCardMode(s.CardMode))
	if s.ActionConditions != nil {
		errorRange = append(errorRange, s.ActionConditions.validate())
	}
	errorRange = append(errorRange, validateDynamicZone(s.Zone))
	if s.ActionExtractors != nil {
		errorRange = append(errorRange, s.ActionExtractors.validate())
	}
	if s.AmountExtractor != nil {
		errorRange = append(errorRange, s.AmountExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (s *entitySelector) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlayerMode(s.PlayerMode))
	if s.PlayerCondition != nil {
		errorRange = append(errorRange, s.PlayerCondition.validate())
	}
	errorRange = append(errorRange, validateCardMode(s.CardMode))
	if s.EntityConditions != nil {
		errorRange = append(errorRange, s.EntityConditions.validate())
	}
	errorRange = append(errorRange, validateDynamicZone(s.Zone))
	if s.EntityExtractors != nil {
		errorRange = append(errorRange, s.EntityExtractors.validate())
	}
	if s.AmountExtractor != nil {
		errorRange = append(errorRange, s.AmountExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (s *placeSelector) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlayerMode(s.PlayerMode))
	if s.PlayerCondition != nil {
		errorRange = append(errorRange, s.PlayerCondition.validate())
	}
	errorRange = append(errorRange, validateCardMode(s.CardMode))
	if s.PlaceConditions != nil {
		errorRange = append(errorRange, s.PlaceConditions.validate())
	}
	errorRange = append(errorRange, validateDynamicZone(s.Zone))
	if s.PlaceExtractors != nil {
		errorRange = append(errorRange, s.PlaceExtractors.validate())
	}
	if s.AmountExtractor != nil {
		errorRange = append(errorRange, s.AmountExtractor.validate())
	}
	return combineErrors(errorRange)
}
