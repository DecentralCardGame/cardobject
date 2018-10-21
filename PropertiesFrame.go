package cardobject

type Property interface {
	IsProperty() bool
}

type CardProperty interface {
	Property
	IsCardProperty() bool
}

type PlayerProperty interface {
	Property
	IsPlayerProperty() bool
}

type IntProperty interface {
	Property
	ExtractIntProp() int
}

type StringProperty interface {
	Property
	ExtractStringProp() string
}

type CardIntProperty interface {
	CardProperty
	ExtractIntProp() int
}

type CardStringProperty interface {
	CardProperty
	ExtractStringProp() string
}

type PlayerIntProperty interface {
	PlayerProperty
	ExtractIntProp() int
}

type PlayerStringProperty interface {
	PlayerProperty
	ExtractStringProp() string
}


type property struct {}
type cardProperty struct {*property}
type playerProperty struct {*property}

type cardIntProperty struct {
	*cardProperty
	cardIntValue int
}

type cardStringProperty struct {
	*cardProperty
	cardStringValue string
}

type playerIntProperty struct {
	*playerProperty
	playerIntValue int
}

type playerStringProperty struct {
	*playerProperty
	playerStringValue string
}


func (p *property) IsProperty() bool {
	return true
}

func (cp *cardProperty) IsCardProperty() bool {
	return true
}

func (cp *playerProperty) IsPlayerProperty() bool {
	return true
}

func (cip *cardIntProperty) ExtractIntProp() int {
	return cip.cardIntValue
}

func (csp *cardStringProperty) ExtractStringProp() string {
	return csp.cardStringValue
}

func (pip *playerIntProperty) ExtractIntProp() int {
	return pip.playerIntValue
}

func (psp *playerStringProperty) ExtractStringProp() string {
	return psp.playerStringValue
}




