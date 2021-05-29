package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type IntValue struct {
	SimpleIntValue *SimpleIntValue  `json:",omitempty"`
	IntVariable    *IntVariableName `json:",omitempty"`
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

func (c ComplexIntValue) InteractionText() string {
	return "§FirstValue §ArithOperator §SecondValue"
}
