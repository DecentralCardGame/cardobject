package cardobject

type Property interface {
	PropertyId() PropertyId
}

type CardProperty interface {
	Property
	CardPropertyId() CardPropertyId
}

type PlayerProperty interface {
	Property
	PlayerPropertyId() PlayerPropertyId
}

type IntProperty interface {
	Property
	IntPropertyId() IntPropertyId
	GetIntVal() int
}

type StringProperty interface {
	Property
	StringPropertyId() StringPropertyId
	GetStringVal() string
}

type intProperty struct {
	val int
}

type stringProperty struct {
	val string
}

type CardIntProperty struct {
	*intProperty
	id CardIntPropertyId
}

type CardStringProperty struct {
	*stringProperty
	id CardStringPropertyId
}

type PlayerIntProperty struct {
	*intProperty
	id PlayerIntPropertyId
}

type PlayerStringProperty struct {
	*stringProperty
	id PlayerStringPropertyId
}


func (ip *intProperty) GetIntVal() int {
	return ip.val
}

func (sp *stringProperty) GetStringVal() string {
	return sp.val
}

func (cip *CardIntProperty) CardIntPropertyId() CardIntPropertyId {
	return cip.id
}

func (cip *CardIntProperty) CardPropertyId() CardPropertyId {
	return cip.CardIntPropertyId()
}

func (cip *CardIntProperty) IntPropertyId() IntPropertyId {
	return cip.CardIntPropertyId()
}

func (cip *CardIntProperty) PropertyId() PropertyId {
	return cip.CardIntPropertyId()
}

func (csp *CardStringProperty) CardStringPropertyId() CardStringPropertyId {
	return csp.id
}

func (csp *CardStringProperty) CardPropertyId() CardPropertyId {
	return csp.CardStringPropertyId()
}

func (csp *CardStringProperty) StringPropertyId() StringPropertyId {
	return csp.CardStringPropertyId()
}

func (csp *CardStringProperty) PropertyId() PropertyId {
	return csp.CardStringPropertyId()
}

func (pip *PlayerIntProperty) PlayerIntPropertyId() PlayerIntPropertyId {
	return pip.id
}

func (pip *PlayerIntProperty) PlayerPropertyId() PlayerPropertyId {
	return pip.PlayerIntPropertyId()
}

func (pip *PlayerIntProperty) IntPropertyId() IntPropertyId {
	return pip.PlayerIntPropertyId()
}

func (pip *PlayerIntProperty) PropertyId() PropertyId {
	return pip.PlayerIntPropertyId()
}

func (psp *PlayerStringProperty) PlayerStringPropertyId() PlayerStringPropertyId {
	return psp.id
}

func (psp *PlayerStringProperty) PlayerPropertyId() PlayerPropertyId {
	return psp.PlayerStringPropertyId()
}

func (psp *PlayerStringProperty) StringPropertyId() StringPropertyId {
	return psp.PlayerStringPropertyId()
}

func (psp *PlayerStringProperty) PropertyId() PropertyId {
	return psp.PlayerStringPropertyId()
}





