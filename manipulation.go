package cardobject

type manipulation interface {
	GetManipulationAttributes() manipulationAttributes
}

type manipulationAttributes struct {
	Duration string
}

type intManipulation struct {
	manipulationAttributes
	Value int
	Operator string
	Property string
	//Selector selector
}

type stringManipulation struct {
	manipulationAttributes
	Value string
	Property string
	//Selector selector
}

func (ma manipulationAttributes) GetManipulationAttributes() manipulationAttributes {
	return ma
}