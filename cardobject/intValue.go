package cardobject

import "github.com/DecentralCardGame/cardobject/jsonschema"

type IntValue struct {
	SimpleIntValue *SimpleIntValue  `json:",omitempty"`
	IntVariable    *IntVariableName `json:",omitempty"`
}

func (i IntValue) Validate(r jsonschema.RootElement) error {
	implementer, err := i.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.Validate(r)
}

func (i IntValue) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(i)
}

type ComplexIntValue struct {
	FirstValue    *IntValue
	SecondValue   *IntValue
	ArithOperator ArithOperator
}

func (c ComplexIntValue) Validate(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(c, r)
}

func (c ComplexIntValue) InteractionText() string {
	return "§FirstValue §ArithOperator §SecondValue"
}
