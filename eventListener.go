package cardobject

import "github.com/DecentralCardGame/jsonschema"

type eventListener struct {
	AttackEventListener       *attackEventListener       `json:",omitempty"`
	BlockEventListener        *blockEventListener        `json:",omitempty"`
	ManipulationEventListener *manipulationEventListener `json:",omitempty"`
	ProductionEventListener   *productionEventListener   `json:",omitempty"`
	TimeEventListener         *timeEventListener         `json:",omitempty"`
	ZoneChangeEventListener   *zoneChangeEventListener   `json:",omitempty"`
}

func (e eventListener) Validate() error {
	return e.ValidateInterface()
}

func (e eventListener) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type attackEventListener struct {
	EntityCondition        *entityCondition  `json:",omitempty"`
	AttackEntityExtractors *entityExtractors `json:",omitempty"`
}

func (a attackEventListener) Validate() error {
	return a.ValidateStruct()
}

func (a attackEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a attackEventListener) GetInteractionText() string {
	return "Whenever an entity §EntityCondition attacks. §AttackEntityExtractors"
}

type blockEventListener struct {
	EntityCondition          *entityCondition  `json:",omitempty"`
	BlockingEntityExtractors *entityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *entityExtractors `json:",omitempty"`
}

func (b blockEventListener) Validate() error {
	return b.ValidateStruct()
}

func (b blockEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b blockEventListener) GetInteractionText() string {
	return "Whenever an entity §EntityCondition blocks. §BlockingEntityExtractors §BlockedEntityExtractors"
}

type manipulationEventListener struct {
	IntManipulationEventListener    *intManipulationEventListener    `json:",omitempty"`
	StringManipulationEventListener *stringManipulationEventListener `json:",omitempty"`
}

func (m manipulationEventListener) Validate() error {
	return m.ValidateInterface()
}

func (m manipulationEventListener) ValidateInterface() error {
	return jsonschema.ValidateInterface(m)
}

type intManipulationEventListener struct {
	IntProperty                cardIntProperty
	IntChangeMode              intChangeMode
	CardCondition              *cardConditions `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors `json:",omitempty"`
	ManipulationValueExtractor *intExtractor   `json:",omitempty"`
}

func (i intManipulationEventListener) Validate() error {
	return i.ValidateStruct()
}

func (i intManipulationEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(i)
}

func (i intManipulationEventListener) GetInteractionText() string {
	return "Whenever §IntProperty on card §CardCondition §IntChangeMode. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type stringManipulationEventListener struct {
	StringProperty             cardStringProperty
	StringChangeMode           stringChangeMode
	CardCondition              *cardConditions  `json:",omitempty"`
	ManipulatedCardExtractor   *cardExtractors  `json:",omitempty"`
	ManipulationValueExtractor *stringExtractor `json:",omitempty"`
}

func (s stringManipulationEventListener) Validate() error {
	return s.ValidateStruct()
}

func (s stringManipulationEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s stringManipulationEventListener) GetInteractionText() string {
	return "Whenever §StringProperty §StringChangeMode on a card §CardCondition. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type productionEventListener struct {
	RessourceTypeCondition    *ressourceCostType `json:",omitempty"`
	ProductionAmountExtractor *intExtractor      `json:",omitempty"`
}

func (p productionEventListener) Validate() error {
	return p.ValidateStruct()
}

func (p productionEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p productionEventListener) GetInteractionText() string {
	return "Whenever one or more §RessourceTypeCondition ressources is produced. §ProductionAmountExtractor"
}

type timeEventListener struct {
	TimeEvent timeEvent
}

func (t timeEventListener) Validate() error {
	return t.ValidateStruct()
}

func (t timeEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t timeEventListener) GetInteractionText() string {
	return "Every §TimeEvent"
}

type zoneChangeEventListener struct {
	Source              dynamicZone
	Destination         zone            `json:",omitempty"`
	CardCondition       *cardConditions `json:",omitempty"`
	MovedCardExtractors *cardExtractors `json:",omitempty"`
}

func (z zoneChangeEventListener) Validate() error {
	return z.ValidateStruct()
}

func (z zoneChangeEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(z)
}

func (z zoneChangeEventListener) GetInteractionText() string {
	return "Whenever a card §CardCondition gets put from §Source to §Destination. §MovedCardExtractors"
}
