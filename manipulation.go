package cardobject

type manipulation interface {
	GetManipulationAttributes() manipulationAttributes
}

type manipulationAttributes struct {
	Duration string
	//Selector selector
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

func (ma manipulationAttributes) GetManipulationAttributes() manipulationAttributes {
	return ma
}