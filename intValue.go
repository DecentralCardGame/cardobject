package cardobject

import "github.com/DecentralCardGame/jsonschema"

type intValue struct {
	*jsonschema.BasicInterface
	ComplexIntValue *complexIntValue `json:",omitempty"`
	SimpleIntValue  *simpleIntValue  `json:",omitempty"`
	IntVariable     *intVariableName `json:",omitempty"`
}

type complexIntValue struct {
	*jsonschema.BasicStruct
	FirstValue    *intValue
	SecondValue   *intValue
	ArithOperator arithOperator
}

func (c complexIntValue) GetInteractionText() string {
	return ""
}
