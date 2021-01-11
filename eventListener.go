package cardobject

import "github.com/DecentralCardGame/jsonschema"

type eventListener struct {
	*jsonschema.BasicInterface
	AttackEventListener       *attackEventListener       `json:",omitempty"`
	BlockEventListener        *blockEventListener        `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ProductionEventListener   *productionEventListener   `json:",omitempty"`
	TimeEventListener         *timeEventListener         `json:",omitempty"`
	ZoneChangeEventListener   *zoneChangeEventListener   `json:",omitempty"`
}

type attackEventListener struct {
	*jsonschema.BasicStruct
	EntityCondition        *entityCondition  `json:",omitempty"`
	AttackEntityExtractors *entityExtractors `json:",omitempty"`
}

func (a attackEventListener) GetInteractionText() string {
	return ""
}

type blockEventListener struct {
	*jsonschema.BasicStruct
	EntityCondition          *entityCondition  `json:",omitempty"`
	BlockingEntityExtractors *entityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *entityExtractors `json:",omitempty"`
}

func (b blockEventListener) GetInteractionText() string {
	return ""
}

type manipulationEventListener struct {
	*jsonschema.BasicInterface
	IntManipulationEventListener    *intManipulationEventListener    `json:",omitempty"`
	StringManipulationEventListener *stringManipulationEventListener `json:",omitempty"`
}

type intManipulationEventListener struct {
	*jsonschema.BasicStruct
	IntProperty                cardIntProperty
	IntChangeMode              intChangeMode
	CardCondition              *cardConditions `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors `json:",omitempty"`
	ManipulationValueExtractor *intExtractor   `json:",omitempty"`
}

func (i intManipulationEventListener) GetInteractionText() string {
	return ""
}

type stringManipulationEventListener struct {
	*jsonschema.BasicStruct
	StringProperty             cardStringProperty
	StringChangeMode           stringChangeMode
	CardCondition              *cardConditions  `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors  `json:",omitempty"`
	ManipulationValueExtractor *stringExtractor `json:",omitempty"`
}

func (s stringManipulationEventListener) GetInteractionText() string {
	return ""
}

type productionEventListener struct {
	*jsonschema.BasicStruct
	RessourceTypeCondition    *ressourceCostType `json:",omitempty"`
	ProductionAmountExtractor *intExtractor      `json:",omitempty"`
}

func (p productionEventListener) GetInteractionText() string {
	return ""
}

type timeEventListener struct {
	*jsonschema.BasicStruct
	TimeEvent timeEvent
}

func (t timeEventListener) GetInteractionText() string {
	return ""
}

type zoneChangeEventListener struct {
	*jsonschema.BasicStruct
	Source              dynamicZone
	Destination         zone            `json:",omitempty"`
	CardCondition       *cardConditions `json:",omitempty"`
	MovedCardExtractors *cardExtractors `json:",omitempty"`
}

func (z zoneChangeEventListener) GetInteractionText() string {
	return ""
}
