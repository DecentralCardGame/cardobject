package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type EventListener struct {
	AttackEventListener       *AttackEventListener       `json:",omitempty"`
	BlockEventListener        *BlockEventListener        `json:",omitempty"`
	ManipulationEventListener *ManipulationEventListener `json:",omitempty"`
	ProductionEventListener   *ProductionEventListener   `json:",omitempty"`
	TimeEventListener         *TimeEventListener         `json:",omitempty"`
	ZoneChangeEventListener   *ZoneChangeEventListener   `json:",omitempty"`
}

func (e EventListener) ValidateType(r jsonschema.RootElement) error {
	implementer, err := e.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (e EventListener) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(e)
}

type AttackEventListener struct {
	EntityCondition        *EntityCondition  `json:",omitempty"`
	AttackEntityExtractors *EntityExtractors `json:",omitempty"`
}

func (a AttackEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a AttackEventListener) InteractionText() string {
	return "Whenever an entity §EntityCondition attacks. §AttackEntityExtractors"
}

type BlockEventListener struct {
	EntityCondition          *EntityCondition  `json:",omitempty"`
	BlockingEntityExtractors *EntityExtractors `json:",omitempty"`
	BlockedEntityExtractors  *EntityExtractors `json:",omitempty"`
}

func (b BlockEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(b, r)
}

func (b BlockEventListener) InteractionText() string {
	return "Whenever an entity §EntityCondition blocks. §BlockingEntityExtractors §BlockedEntityExtractors"
}

type ManipulationEventListener struct {
	IntManipulationEventListener    *IntManipulationEventListener    `json:",omitempty"`
	StringManipulationEventListener *StringManipulationEventListener `json:",omitempty"`
}

func (m ManipulationEventListener) ValidateType(r jsonschema.RootElement) error {
	implementer, err := m.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (m ManipulationEventListener) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(m)
}

type IntManipulationEventListener struct {
	IntProperty                CardIntProperty
	IntChangeMode              IntChangeMode
	CardCondition              *CardConditions `json:",omitempty"`
	ManipulatedCardExtractor   *CardExtractors `json:",omitempty"`
	ManipulationValueExtractor *IntExtractor   `json:",omitempty"`
}

func (i IntManipulationEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(i, r)
}

func (i IntManipulationEventListener) InteractionText() string {
	return "Whenever §IntProperty on card §CardCondition §IntChangeMode. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type StringManipulationEventListener struct {
	StringProperty             CardStringProperty
	StringChangeMode           StringChangeMode
	CardCondition              *CardConditions  `json:",omitempty"`
	ManipulatedCardExtractor   *CardExtractors  `json:",omitempty"`
	ManipulationValueExtractor *StringExtractor `json:",omitempty"`
}

func (s StringManipulationEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s StringManipulationEventListener) InteractionText() string {
	return "Whenever §StringProperty §StringChangeMode on a card §CardCondition. §ManipulatedCardExtractors §ManipulationValueExtractor"
}

type ProductionEventListener struct {
	ClassCondition            *Class        `json:",omitempty"`
	ProductionAmountExtractor *IntExtractor `json:",omitempty"`
}

func (p ProductionEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p ProductionEventListener) InteractionText() string {
	return "Whenever one or more §ClassCondition mana is produced. §ProductionAmountExtractor"
}

type TimeEventListener struct {
	TimeEvent TimeEvent
}

func (t TimeEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(t, r)
}

func (t TimeEventListener) InteractionText() string {
	return "Every §TimeEvent"
}

type ZoneChangeEventListener struct {
	Source              *DynamicZone    `json:",omitempty"`
	Destination         *Zone           `json:",omitempty"`
	CardCondition       *CardConditions `json:",omitempty"`
	MovedCardExtractors *CardExtractors `json:",omitempty"`
}

func (z ZoneChangeEventListener) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(z, r)
}

func (z ZoneChangeEventListener) InteractionText() string {
	return "Whenever a card §CardCondition gets put from §Source to §Destination. §MovedCardExtractors"
}
