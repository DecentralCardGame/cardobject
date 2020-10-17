package cardobject

import "errors"

type intValue struct {
	ComplexIntValue *complexIntValue
	SimpleIntValue  *int
	IntVariable     *string
}

type complexIntValue struct {
	FirstValue    *complexIntValue
	SecondValue   *complexIntValue
	ArithOperator string
}

func (v *intValue) validate() error {
	if v.ComplexIntValue != nil {
		if v.SimpleIntValue != nil {
			return errors.New("IntValue implemented by not exactly one option")
		}
		if v.IntVariable != nil {
			return errors.New("IntValue implemented by not exactly one option")
		}
		return v.ComplexIntValue.validate()
	}
	if v.SimpleIntValue != nil {
		if v.IntVariable != nil {
			return errors.New("IntValue implemented by not exactly one option")
		}
		return validateSimpleInt(*v.SimpleIntValue)
	}
	return validateIntVariableName(*v.IntVariable)
}

func (v *complexIntValue) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, v.FirstValue.validate())
	errorRange = append(errorRange, v.SecondValue.validate())
	errorRange = append(errorRange, validateArithOperator(v.ArithOperator))
	return combineErrors(errorRange)
}
