package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var abilityEffectOperators []string = []string{"GAIN", "LOSE"}
var intOperators []string = []string{"SET", "ADD", "SUBTRACT"}
var stringOperators []string = []string{"SET"}

var arithOperators []string = []string{"ADD", "SUBTRACT"}

type abilityEffectOperator jsonschema.BasicEnum

func (a abilityEffectOperator) Validate() error {
	return a.ValidateEnum()
}

func (a abilityEffectOperator) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a abilityEffectOperator) GetEnumValues() []string {
	return abilityEffectOperators
}

type intOperator jsonschema.BasicEnum

func (i intOperator) Validate() error {
	return i.ValidateEnum()
}

func (i intOperator) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i intOperator) GetEnumValues() []string {
	return intOperators
}

type stringOperator jsonschema.BasicEnum

func (s stringOperator) Validate() error {
	return s.ValidateEnum()
}

func (s stringOperator) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s stringOperator) GetEnumValues() []string {
	return stringOperators
}

type arithOperator jsonschema.BasicEnum

func (a arithOperator) Validate() error {
	return a.ValidateEnum()
}

func (a arithOperator) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a arithOperator) GetEnumValues() []string {
	return arithOperators
}
