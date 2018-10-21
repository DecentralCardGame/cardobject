package cardobject

type Effect struct {
	zoneChanges []ZoneChange
	manipulations []Manipulation
	production []Ressource
}

type Manipulation interface {
	GetSelector() Selector
	GetManipulationValue() interface{}
	GetTargetProperty() Property
}

type IntManipulation interface {
	Manipulation
	GetIntManipulationValue() int
	GetTargetIntProperty() IntProperty
}

type StringManipulation interface {
	Manipulation
	GetStringManipulationValue() string
	GetTargetStringProperty() StringProperty
}

func NewIntManipulation(s Selector, v int, p IntProperty) IntManipulation {
	return &intManipulation{&manipulation{s}, v, p}
}

func NewStringManipulation(s Selector, v string, p StringProperty) StringManipulation {
	return &stringManipulation{&manipulation{s}, v, p}
}


type manipulation struct {
	selector Selector
}

type intManipulation struct {
	*manipulation
	value int
	property IntProperty
}

type stringManipulation struct {
	*manipulation
	value string
	property StringProperty
}


func (m *manipulation) GetSelector() Selector {
	return m.selector
}

func (im *intManipulation) GetManipulationValue() interface{} {
	return im.GetIntManipulationValue()
}

func (sm *stringManipulation) GetManipulationValue() interface{} {
	return sm.GetStringManipulationValue()
}

func (im *intManipulation) GetTargetProperty() Property {
	return im.GetTargetIntProperty()
}

func (sm *stringManipulation) GetTargetProperty() Property {
	return sm.GetTargetStringProperty()
}

func (im *intManipulation) GetIntManipulationValue() int {
	return im.value
}

func (sm *stringManipulation) GetStringManipulationValue() string {
	return sm.value
}

func (im *intManipulation) GetTargetIntProperty() IntProperty {
	return im.property
}

func (sm *stringManipulation) GetTargetStringProperty() StringProperty {
	return sm.property
}

