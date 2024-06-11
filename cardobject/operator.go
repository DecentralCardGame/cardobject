package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

// Add Operator
const Add = "ADD"

// Gain Operator
const Gain = "GAIN"

// Lose Operator
const Lose = "LOSE"

// Set Operator
const Set = "SET"

// Subtract Operator
const Subtract = "SUBTRACT"

var abilityEffectOperators []string = []string{Gain}
var intOperators []string = []string{Set, Add, Subtract}
var stringOperators []string = []string{Set}

var arithOperators []string = []string{Add, Subtract}

type AbilityEffectOperator jsonschema.BasicEnum

func (a AbilityEffectOperator) ValidateType(r jsonschema.RootElement) error {
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

func (i IntOperator) ValidateType(r jsonschema.RootElement) error {
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

func (s StringOperator) ValidateType(r jsonschema.RootElement) error {
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

func (a ArithOperator) ValidateType(r jsonschema.RootElement) error {
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
