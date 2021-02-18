package cardobject

import "github.com/DecentralCardGame/jsonschema"

type intValue struct {
	ComplexIntValue *complexIntValue `json:",omitempty"`
	SimpleIntValue  *simpleIntValue  `json:",omitempty"`
	IntVariable     *intVariableName `json:",omitempty"`
}

func (i intValue) Validate() error {
	return i.ValidateInterface()
}

func (i intValue) ValidateInterface() error {
	return jsonschema.ValidateInterface(i)
}

type complexIntValue struct {
	FirstValue    *intValue
	SecondValue   *intValue
	ArithOperator arithOperator
}

func (c complexIntValue) Validate() error {
	return c.ValidateStruct()
}

func (c complexIntValue) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c complexIntValue) GetInteractionText() string {
	return "§FirstValue §ArithOperator §SecondValue"
}
