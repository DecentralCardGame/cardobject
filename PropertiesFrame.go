package cardobject

type Property interface {
	IsProperty() bool
}

type CardProperty interface {
	IsCardProperty() bool
}

type PlayerProperty interface {
	IsPlayerProperty() bool
}

type IntProperty interface {
	ExtractIntProp() int
}

type StringProperty interface {
	ExtractStringProp() string
}

type CardIntProperty interface {
	IsCardProperty() bool
	ExtractIntProp() int
}

type CardStringProperty interface {
	IsCardProperty() bool
	ExtractStringProp() string
}

type PlayerIntProperty interface {
	IsPlayerProperty() bool
	ExtractIntProp() int
}

type PlayerStringProperty interface {
	IsPlayerProperty() bool
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




