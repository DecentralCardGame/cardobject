package cardobject

type PropertyId interface {
	IsPropertyId() bool
}

type CardPropertyId interface {
	PropertyId
	IsCardPropertyId() bool
}

type PlayerPropertyId interface {
	PropertyId
	IsPlayerPropertyId() bool
}

type IntPropertyId interface {
	PropertyId
	IsIntPropertyId() bool
}

type StringPropertyId interface {
	PropertyId
	IsStringPropertyId() bool
}

type CardIntPropertyId int
type CardStringPropertyId int
type PlayerIntPropertyId int
type PlayerStringPropertyId int


func (cipi CardIntPropertyId) IsCardPropertyId() bool {
	return true
}

func (cipi CardIntPropertyId) IsIntPropertyId() bool {
	return true
}

func (cspi CardStringPropertyId) IsCardPropertyId() bool {
	return true
}

func (cspi CardStringPropertyId) IsStringPropertyId() bool {
	return true
}

func (pipi PlayerIntPropertyId) IsPlayerPropertyId() bool {
	return true
}

func (pipi PlayerIntPropertyId) IsIntPropertyId() bool {
	return true
}

func (pspi PlayerStringPropertyId) IsPlayerPropertyId() bool {
	return true
}

func (pspi PlayerStringPropertyId) IsStringPropertyId() bool {
	return true
}

func (cipi CardIntPropertyId) IsPropertyId() bool {
	return true
}

func (cipi CardStringPropertyId) IsPropertyId() bool {
	return true
}

func (cipi PlayerIntPropertyId) IsPropertyId() bool {
	return true
}

func (cipi PlayerStringPropertyId) IsPropertyId() bool {
	return true
}