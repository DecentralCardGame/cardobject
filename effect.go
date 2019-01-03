package cardobject

type Effect struct {
	zoneChanges []ZoneChange
	manipulations []Manipulation
	production []Ressource
}

type Manipulation interface {
	GetCardSelector() CardSelector
	GetManipulation() interface{}
	GetTargetPropertyId() PropertyId
}

func NewIntManipulation(cs CardSelector, v IntInserter, p IntPropertyId) *IntManipulation {
	return &IntManipulation{&manipulation{cs}, v, p}
}

func NewStringManipulation(cs CardSelector, v StringInserter, p StringPropertyId) *StringManipulation {
	return &StringManipulation{&manipulation{cs}, v, p}
}


type manipulation struct {
	selector CardSelector
}

type IntManipulation struct {
	*manipulation
	val IntInserter
	prop IntPropertyId
}

type StringManipulation struct {
	*manipulation
	val StringInserter
	prop StringPropertyId
}


func (m *manipulation) GetCardSelector() CardSelector {
	return m.selector
}

func (im *IntManipulation) GetManipulation() interface{} {
	return im.GetIntManipulation()
}

func (sm *StringManipulation) GetManipulation() interface{} {
	return sm.GetStringManipulation()
}

func (im *IntManipulation) GetTargetPropertyId() PropertyId {
	return im.GetTargetIntPropertyId()
}

func (sm *StringManipulation) GetTargetPropertyId() PropertyId {
	return sm.GetTargetStringPropertyId()
}

func (im *IntManipulation) GetIntManipulation() IntInserter {
	return im.val
}

func (sm *StringManipulation) GetStringManipulation() StringInserter {
	return sm.val
}

func (im *IntManipulation) GetTargetIntPropertyId() IntPropertyId {
	return im.prop
}

func (sm *StringManipulation) GetTargetStringPropertyId() StringPropertyId {
	return sm.prop
}

