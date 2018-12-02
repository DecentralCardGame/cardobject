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

type CardIntProperty interface {
	IntProperty
	CardPropertyId() CardPropertyId
	CardIntPropertyId() CardIntPropertyId
}

type CardStringProperty interface {
	StringProperty
	CardPropertyId() CardPropertyId
	CardStringPropertyId() CardStringPropertyId
}

type PlayerIntProperty interface {
	IntProperty
	PlayerPropertyId() PlayerPropertyId
	PlayerIntPropertyId() PlayerIntPropertyId
}

type PlayerStringProperty interface {
	StringProperty
	PlayerPropertyId() PlayerPropertyId
	PlayerStringPropertyId() PlayerStringPropertyId
}


type intProperty struct {
	val int
}

type stringProperty struct {
	val string
}

type cardIntProperty struct {
	*intProperty
	id CardIntPropertyId
}

type cardStringProperty struct {
	*stringProperty
	id CardStringPropertyId
}

type playerIntProperty struct {
	*intProperty
	id PlayerIntPropertyId
}

type playerStringProperty struct {
	*stringProperty
	id PlayerStringPropertyId
}


func (ip *intProperty) GetIntVal() int {
	return ip.val
}

func (sp *stringProperty) GetStringVal() string {
	return sp.val
}

func (cip *cardIntProperty) CardIntPropertyId() CardIntPropertyId {
	return cip.id
}

func (cip *cardIntProperty) CardPropertyId() CardPropertyId {
	return cip.CardIntPropertyId()
}

func (cip *cardIntProperty) IntPropertyId() IntPropertyId {
	return cip.CardIntPropertyId()
}

func (cip *cardIntProperty) PropertyId() PropertyId {
	return cip.CardIntPropertyId()
}

func (csp *cardStringProperty) CardStringPropertyId() CardStringPropertyId {
	return csp.id
}

func (csp *cardStringProperty) CardPropertyId() CardPropertyId {
	return csp.CardStringPropertyId()
}

func (csp *cardStringProperty) StringPropertyId() StringPropertyId {
	return csp.CardStringPropertyId()
}

func (csp *cardStringProperty) PropertyId() PropertyId {
	return csp.CardStringPropertyId()
}

func (pip *playerIntProperty) PlayerIntPropertyId() PlayerIntPropertyId {
	return pip.id
}

func (pip *playerIntProperty) PlayerPropertyId() PlayerPropertyId {
	return pip.PlayerIntPropertyId()
}

func (pip *playerIntProperty) IntPropertyId() IntPropertyId {
	return pip.PlayerIntPropertyId()
}

func (pip *playerIntProperty) PropertyId() PropertyId {
	return pip.PlayerIntPropertyId()
}

func (psp *playerStringProperty) PlayerStringPropertyId() PlayerStringPropertyId {
	return psp.id
}

func (psp *playerStringProperty) PlayerPropertyId() PlayerPropertyId {
	return psp.PlayerStringPropertyId()
}

func (psp *playerStringProperty) StringPropertyId() StringPropertyId {
	return psp.PlayerStringPropertyId()
}

func (psp *playerStringProperty) PropertyId() PropertyId {
	return psp.PlayerStringPropertyId()
}





