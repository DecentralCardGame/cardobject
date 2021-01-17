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
	return "Whenever an entity §EntityCondition attacks. §AttackEntityExtractors"
}

type blockEventListener struct {
	*jsonschema.BasicStruct
	EntityCondition          *entityCondition  `json:",omitempty"`
	BlockingEntityExtractors *entityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *entityExtractors `json:",omitempty"`
}

func (b blockEventListener) GetInteractionText() string {
	return "Whenever an entity §EntityCondition blocks. §BlockingEntityExtractors §BlockedEntityExtractors"
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
	return "Whenever §IntProperty on card §CardCondition §IntChangeMode. §ManipulatedCardExtractors §ManipulationValueExtractor"
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
	return "Whenever §StringProperty §StringChangeMode on a card §CardCondition. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type productionEventListener struct {
	*jsonschema.BasicStruct
	RessourceTypeCondition    *ressourceCostType `json:",omitempty"`
	ProductionAmountExtractor *intExtractor      `json:",omitempty"`
}

func (p productionEventListener) GetInteractionText() string {
	return "Whenever one or more §RessourceTypeCondition ressources is produced. §ProductionAmountExtractor"
}

type timeEventListener struct {
	*jsonschema.BasicStruct
	TimeEvent timeEvent
}

func (t timeEventListener) GetInteractionText() string {
	return "Every §TimeEvent"
}

type zoneChangeEventListener struct {
	*jsonschema.BasicStruct
	Source              dynamicZone
	Destination         zone            `json:",omitempty"`
	CardCondition       *cardConditions `json:",omitempty"`
	MovedCardExtractors *cardExtractors `json:",omitempty"`
}

func (z zoneChangeEventListener) GetInteractionText() string {
	return "Whenever a card §CardCondition gets put from §Source to §Destination. §MovedCardExtractors"
}
