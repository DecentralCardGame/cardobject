package cardobject

type Effect struct {
	zoneChanges []ZoneChange
	manipulations []Manipulation
	production []Ressource
}

type Manipulation interface {
	GetSelector() Selector
	GetManipulation() interface{}
	GetTargetPropertyId() PropertyId
}

type IntManipulation interface {
	Manipulation
	GetIntManipulation() IntInserter
	GetTargetIntPropertyId() IntPropertyId
}

type StringManipulation interface {
	Manipulation
	GetStringManipulation() StringInserter
	GetTargetStringPropertyId() StringPropertyId
}

func NewIntManipulation(s Selector, v IntInserter, p IntPropertyId) IntManipulation {
	return &intManipulation{&manipulation{s}, v, p}
}

func NewStringManipulation(s Selector, v StringInserter, p StringPropertyId) StringManipulation {
	return &stringManipulation{&manipulation{s}, v, p}
}


type manipulation struct {
	selector Selector
}

type intManipulation struct {
	*manipulation
	val IntInserter
	prop IntPropertyId
}

type stringManipulation struct {
	*manipulation
	val StringInserter
	prop StringPropertyId
}


func (m *manipulation) GetSelector() Selector {
	return m.selector
}

func (im *intManipulation) GetManipulation() interface{} {
	return im.GetIntManipulation()
}

func (sm *stringManipulation) GetManipulation() interface{} {
	return sm.GetStringManipulation()
}

func (im *intManipulation) GetTargetPropertyId() PropertyId {
	return im.GetTargetIntPropertyId()
}

func (sm *stringManipulation) GetTargetPropertyId() PropertyId {
	return sm.GetTargetStringPropertyId()
}

func (im *intManipulation) GetIntManipulation() IntInserter {
	return im.val
}

func (sm *stringManipulation) GetStringManipulation() StringInserter {
	return sm.val
}

func (im *intManipulation) GetTargetIntPropertyId() IntPropertyId {
	return im.prop
}

func (sm *stringManipulation) GetTargetStringPropertyId() StringPropertyId {
	return sm.prop
}

