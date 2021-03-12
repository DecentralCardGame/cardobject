package cardobject

import "github.com/DecentralCardGame/jsonschema"

type IntValue struct {
	ComplexIntValue *ComplexIntValue `json:",omitempty"`
	SimpleIntValue  *SimpleIntValue  `json:",omitempty"`
	IntVariable     *IntVariableName `json:",omitempty"`
}

func (i IntValue) Validate() error {
	return i.ValidateInterface()
}

func (i IntValue) ValidateInterface() error {
	return jsonschema.ValidateInterface(i)
}

type ComplexIntValue struct {
	FirstValue    *IntValue
	SecondValue   *IntValue
	ArithOperator ArithOperator
}

func (c ComplexIntValue) Validate() error {
	return c.ValidateStruct()
}

func (c ComplexIntValue) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c ComplexIntValue) GetInteractionText() string {
	return "§FirstValue §ArithOperator §SecondValue"
}
