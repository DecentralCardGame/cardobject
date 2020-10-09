package cardobject

import "errors"

type eventListener struct {
	AttackEventListener       *attackEventListener       `json:",omitempty"`
	BlockEventListener        *blockEventListener        `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ProductionEventListener   *productionEventListener   `json:",omitempty"`
	TimeEventListener         *timeEventListener         `json:",omitempty"`
	ZoneChangeEventListener   *zoneChangeEventListener   `json:",omitempty"`
}

type attackEventListener struct {
	EntityCondition        *entityCondition  `json:",omitempty"`
	AttackEntityExtractors *entityExtractors `json:",omitempty"`
}

type blockEventListener struct {
	EntityCondition          *entityCondition  `json:",omitempty"`
	BlockingEntityExtractors *entityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *entityExtractors `json:",omitempty"`
}

type manipulationEventListener struct {
	IntManipulationEventListener    *intManipulationEventListener    `json:",omitempty"`
	StringManipulationEventListener *stringManipulationEventListener `json:",omitempty"`
}

type intManipulationEventListener struct {
	Property                   string
	IntChangeMode              string
	CardCondition              *cardConditions `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors `json:",omitempty"`
	ManipulationValueExtractor *intExtractor   `json:",omitempty"`
}

type stringManipulationEventListener struct {
	Property                   string
	StringChangeMode           string
	CardCondition              *cardConditions  `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors  `json:",omitempty"`
	ManipulationValueExtractor *stringExtractor `json:",omitempty"`
}

type productionEventListener struct {
	RessourceTypeCondition    *ressourceCostType `json:",omitempty"`
	ProductionAmountExtractor *intExtractor      `json:",omitempty"`
}

type timeEventListener struct {
	TimeEvent string
}

type zoneChangeEventListener struct {
	Source              string
	Destination         string          `json:",omitempty"`
	CardCondition       *cardConditions `json:",omitempty"`
	MovedCardExtractors *cardExtractors `json:",omitempty"`
}

func (e *eventListener) validate() error {
	possibleImplementer := []validateable{e.AttackEventListener, e.BlockEventListener, e.ManipulationEventListener, e.ProductionEventListener, e.TimeEventListener, e.ZoneChangeEventListener}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("EventListener implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *attackEventListener) validate() error {
	errorRange := []error{}
	if e.EntityCondition != nil {
		errorRange = append(errorRange, e.EntityCondition.validate())
	}
	if e.AttackEntityExtractors != nil {
		errorRange = append(errorRange, e.AttackEntityExtractors.validate())
	}
	return combineErrors(errorRange)
}

func (e *blockEventListener) validate() error {
	errorRange := []error{}
	if e.EntityCondition != nil {
		errorRange = append(errorRange, e.EntityCondition.validate())
	}
	if e.BlockedEntityExtractors != nil {
		errorRange = append(errorRange, e.BlockedEntityExtractors.validate())
	}
	if e.BlockingEntityExtractors != nil {
		errorRange = append(errorRange, e.BlockingEntityExtractors.validate())
	}
	return combineErrors(errorRange)
}

func (e *manipulationEventListener) validate() error {
	possibleImplementer := []validateable{e.IntManipulationEventListener, e.StringManipulationEventListener}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("ManipulationEventListener implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *intManipulationEventListener) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardIntProperty(e.Property))
	errorRange = append(errorRange, validateIntChangeMode(e.IntChangeMode))
	if e.CardCondition != nil {
		errorRange = append(errorRange, e.CardCondition.validate())
	}
	if e.ManipulatedCardExtractor != nil {
		errorRange = append(errorRange, e.ManipulatedCardExtractor.validate())
	}
	if e.ManipulationValueExtractor != nil {
		errorRange = append(errorRange, e.ManipulationValueExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e *stringManipulationEventListener) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardStringProperty(e.Property))
	errorRange = append(errorRange, validateStringChangeMode(e.StringChangeMode))
	if e.CardCondition != nil {
		errorRange = append(errorRange, e.CardCondition.validate())
	}
	if e.ManipulatedCardExtractor != nil {
		errorRange = append(errorRange, e.ManipulatedCardExtractor.validate())
	}
	if e.ManipulationValueExtractor != nil {
		errorRange = append(errorRange, e.ManipulationValueExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e *productionEventListener) validate() error {
	errorRange := []error{}
	if e.ProductionAmountExtractor != nil {
		errorRange = append(errorRange, e.ProductionAmountExtractor.validate())
	}
	return combineErrors(errorRange)
}

func (e *timeEventListener) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateTimeEvent(e.TimeEvent))
	return combineErrors(errorRange)
}

func (e *zoneChangeEventListener) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateDynamicZone(e.Source))
	errorRange = append(errorRange, validateZone(e.Destination))

	if e.CardCondition != nil {
		errorRange = append(errorRange, e.CardCondition.validate())
	}
	if e.MovedCardExtractors != nil {
		errorRange = append(errorRange, e.MovedCardExtractors.validate())
	}
	return combineErrors(errorRange)
}
