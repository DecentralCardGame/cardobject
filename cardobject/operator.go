package cardobject

import (
	"errors"
	"strings"

	"cardobject/jsonschema"
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
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("AbilityEffectOperator must be one of: " + strings.Join(abilityEffectOperators, ","))
}

func (a AbilityEffectOperator) EnumValues() []string {
	return abilityEffectOperators
}

type IntOperator jsonschema.BasicEnum

func (i IntOperator) Validate() error {
	return i.ValidateEnum()
}

func (i IntOperator) ValidateEnum() error {
	values := i.EnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("IntOperator must be one of: " + strings.Join(intOperators, ","))
}

func (i IntOperator) EnumValues() []string {
	return intOperators
}

type StringOperator jsonschema.BasicEnum

func (s StringOperator) Validate() error {
	return s.ValidateEnum()
}

func (s StringOperator) ValidateEnum() error {
	values := s.EnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("StringOperator must be one of: " + strings.Join(stringComparators, ","))
}

func (s StringOperator) EnumValues() []string {
	return stringOperators
}

type ArithOperator jsonschema.BasicEnum

func (a ArithOperator) Validate() error {
	return a.ValidateEnum()
}

func (a ArithOperator) ValidateEnum() error {
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ArithOperator must be one of: " + strings.Join(arithOperators, ","))
}

func (a ArithOperator) EnumValues() []string {
	return arithOperators
}
