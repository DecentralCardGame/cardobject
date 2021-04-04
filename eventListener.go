package cardobject

import "github.com/DecentralCardGame/jsonschema"

type EventListener struct {
	AttackEventListener       *AttackEventListener       `json:",omitempty"`
	BlockEventListener        *BlockEventListener        `json:",omitempty"`
	ManipulationEventListener *ManipulationEventListener `json:",omitempty"`
	ProductionEventListener   *ProductionEventListener   `json:",omitempty"`
	TimeEventListener         *TimeEventListener         `json:",omitempty"`
	ZoneChangeEventListener   *ZoneChangeEventListener   `json:",omitempty"`
}

func (e EventListener) Validate() error {
	return e.ValidateInterface()
}

func (e EventListener) ValidateInterface() error {
	return jsonschema.ValidateInterface(e)
}

type AttackEventListener struct {
	EntityCondition        *EntityCondition  `json:",omitempty"`
	AttackEntityExtractors *EntityExtractors `json:",omitempty"`
}

func (a AttackEventListener) Validate() error {
	return a.ValidateStruct()
}

func (a AttackEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a AttackEventListener) GetInteractionText() string {
	return "Whenever an entity §EntityCondition attacks. §AttackEntityExtractors"
}

type BlockEventListener struct {
	EntityCondition          *EntityCondition  `json:",omitempty"`
	BlockingEntityExtractors *EntityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *EntityExtractors `json:",omitempty"`
}

func (b BlockEventListener) Validate() error {
	return b.ValidateStruct()
}

func (b BlockEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(b)
}

func (b BlockEventListener) GetInteractionText() string {
	return "Whenever an entity §EntityCondition blocks. §BlockingEntityExtractors §BlockedEntityExtractors"
}

type ManipulationEventListener struct {
	IntManipulationEventListener    *IntManipulationEventListener    `json:",omitempty"`
	StringManipulationEventListener *StringManipulationEventListener `json:",omitempty"`
}

func (m ManipulationEventListener) Validate() error {
	return m.ValidateInterface()
}

func (m ManipulationEventListener) ValidateInterface() error {
	return jsonschema.ValidateInterface(m)
}

type IntManipulationEventListener struct {
	IntProperty                CardIntProperty
	IntChangeMode              IntChangeMode
	CardCondition              *CardConditions `json:",omitempty"`
	ManipulatedCardExtractor   *CardExtractors `json:",omitempty"`
	ManipulationValueExtractor *IntExtractor   `json:",omitempty"`
}

func (i IntManipulationEventListener) Validate() error {
	return i.ValidateStruct()
}

func (i IntManipulationEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(i)
}

func (i IntManipulationEventListener) GetInteractionText() string {
	return "Whenever §IntProperty on card §CardCondition §IntChangeMode. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type StringManipulationEventListener struct {
	StringProperty             CardStringProperty
	StringChangeMode           StringChangeMode
	CardCondition              *CardConditions  `json:",omitempty"`
	ManipulatedCardExtractor   *CardExtractors  `json:",omitempty"`
	ManipulationValueExtractor *StringExtractor `json:",omitempty"`
}

func (s StringManipulationEventListener) Validate() error {
	return s.ValidateStruct()
}

func (s StringManipulationEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s StringManipulationEventListener) GetInteractionText() string {
	return "Whenever §StringProperty §StringChangeMode on a card §CardCondition. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type ProductionEventListener struct {
	RessourceTypeCondition    *RessourceCostType `json:",omitempty"`
	ProductionAmountExtractor *IntExtractor      `json:",omitempty"`
}

func (p ProductionEventListener) Validate() error {
	return p.ValidateStruct()
}

func (p ProductionEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p ProductionEventListener) GetInteractionText() string {
	return "Whenever one or more §RessourceTypeCondition ressources is produced. §ProductionAmountExtractor"
}

type TimeEventListener struct {
	TimeEvent TimeEvent
}

func (t TimeEventListener) Validate() error {
	return t.ValidateStruct()
}

func (t TimeEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(t)
}

func (t TimeEventListener) GetInteractionText() string {
	return "Every §TimeEvent"
}

type ZoneChangeEventListener struct {
	Source              *DynamicZone    `json:",omitempty"`
	Destination         *Zone           `json:",omitempty"`
	CardCondition       *CardConditions `json:",omitempty"`
	MovedCardExtractors *CardExtractors `json:",omitempty"`
}

func (z ZoneChangeEventListener) Validate() error {
	return z.ValidateStruct()
}

func (z ZoneChangeEventListener) ValidateStruct() error {
	return jsonschema.ValidateStruct(z)
}

func (z ZoneChangeEventListener) GetInteractionText() string {
	return "Whenever a card §CardCondition gets put from §Source to §Destination. §MovedCardExtractors"
}
