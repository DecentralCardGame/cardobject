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

type CardIntPropertyId interface {
	CardPropertyId
	IsIntPropertyId() bool
}

type CardStringPropertyId interface {
	CardPropertyId
	IsStringPropertyId() bool
}

type PlayerIntPropertyId interface {
	PlayerPropertyId
	IsIntPropertyId() bool
}

type PlayerStringPropertyId interface {
	PlayerPropertyId
	IsStringPropertyId() bool
}

type cardIntPropertyId int
type cardStringPropertyId int
type playerIntPropertyId int
type playerStringPropertyId int


func (cipi cardIntPropertyId) IsCardPropertyId() bool {
	return true
}

func (cipi cardIntPropertyId) IsIntPropertyId() bool {
	return true
}

func(cspi cardStringPropertyId) IsCardPropertyId() bool {
	return true
}

func(cspi cardStringPropertyId) IsStringPropertyId() bool {
	return true
}

func (pipi playerIntPropertyId) IsPlayerPropertyId() bool {
	return true
}

func (pipi playerIntPropertyId) IsIntPropertyId() bool {
	return true
}

func (pspi playerStringPropertyId) IsPlayerPropertyId() bool {
	return true
}

func (pspi playerStringPropertyId) IsStringPropertyId() bool {
	return true
}

func (cipi cardIntPropertyId) IsPropertyId() bool {
	return true
}

func (cipi cardStringPropertyId) IsPropertyId() bool {
	return true
}

func (cipi playerIntPropertyId) IsPropertyId() bool {
	return true
}

func (cipi playerStringPropertyId) IsPropertyId() bool {
	return true
}