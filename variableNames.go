package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var intVariableNames []string = []string{"X", "Y", "Z"}
var stringVariableNames []string = []string{"A", "B", "C"}
var targetVariableNames []string = []string{"M", "T"}

type IntVariableName jsonschema.BasicEnum

func (i IntVariableName) Validate() error {
	return i.ValidateEnum()
}

func (i IntVariableName) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i IntVariableName) GetEnumValues() []string {
	return intVariableNames
}

type StringVariableName jsonschema.BasicEnum

func (s StringVariableName) Validate() error {
	return s.ValidateEnum()
}

func (s StringVariableName) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s StringVariableName) GetEnumValues() []string {
	return stringVariableNames
}

type TargetVariableName jsonschema.BasicEnum

func (t TargetVariableName) Validate() error {
	return t.ValidateEnum()
}

func (t TargetVariableName) ValidateEnum() error {
	values := t.GetEnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("")
}

func (t TargetVariableName) GetEnumValues() []string {
	return targetVariableNames
}
