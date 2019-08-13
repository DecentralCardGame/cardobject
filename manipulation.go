package cardobject

type manipulation interface {
	GetManipulationAttributes() manipulationAttributes
}

type manipulationAttributes struct {
	Duration string
	Selector selectorWrapper
}

type intManipulation struct {
	manipulationAttributes
	Value int
	Operator string
	Property string
}

type stringManipulation struct {
	manipulationAttributes
	Value string
	Property string
}

type manipulationWrapper struct {
	IntManipulation *intManipulation `json:",omitempty"`
	StringManipulation *stringManipulation `json:",omitempty"`
}

func (ma manipulationAttributes) GetManipulationAttributes() manipulationAttributes {
	return ma
}