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
	GetIntValue() int
}

type StringProperty interface {
	GetStringValue() string
}

type CardIntProperty interface {
	IsCardProperty() bool
	GetIntValue() int
}

type CardStringProperty interface {
	IsCardProperty() bool
	GetStringValue() string
}

type PlayerIntProperty interface {
	IsPlayerProperty() bool
	GetIntValue() int
}

type PlayerStringProperty interface {
	IsPlayerProperty() bool
	GetStringValue() string
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




