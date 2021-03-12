package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var abilityEffectOperators []string = []string{"GAIN", "LOSE"}
var intOperators []string = []string{"SET", "ADD", "SUBTRACT"}
var stringOperators []string = []string{"SET"}

var arithOperators []string = []string{"ADD", "SUBTRACT"}

type AbilityEffectOperator jsonschema.BasicEnum

func (a AbilityEffectOperator) Validate() error {
	return a.ValidateEnum()
}

func (a AbilityEffectOperator) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a AbilityEffectOperator) GetEnumValues() []string {
	return abilityEffectOperators
}

type IntOperator jsonschema.BasicEnum

func (i IntOperator) Validate() error {
	return i.ValidateEnum()
}

func (i IntOperator) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i IntOperator) GetEnumValues() []string {
	return intOperators
}

type StringOperator jsonschema.BasicEnum

func (s StringOperator) Validate() error {
	return s.ValidateEnum()
}

func (s StringOperator) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s StringOperator) GetEnumValues() []string {
	return stringOperators
}

type ArithOperator jsonschema.BasicEnum

func (a ArithOperator) Validate() error {
	return a.ValidateEnum()
}

func (a ArithOperator) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a ArithOperator) GetEnumValues() []string {
	return arithOperators
}
