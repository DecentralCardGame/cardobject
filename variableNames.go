package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var intVariableNames []string = []string{"X", "Y", "Z"}
var stringVariableNames []string = []string{"A", "B", "C"}
var targetVariableNames []string = []string{"M", "T"}

type intVariableName jsonschema.BasicEnum

func (i intVariableName) Validate() error {
	return i.ValidateEnum()
}

func (i intVariableName) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i intVariableName) GetEnumValues() []string {
	return intVariableNames
}

type stringVariableName jsonschema.BasicEnum

func (s stringVariableName) Validate() error {
	return s.ValidateEnum()
}

func (s stringVariableName) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s stringVariableName) GetEnumValues() []string {
	return stringVariableNames
}

type targetVariableName jsonschema.BasicEnum

func (t targetVariableName) Validate() error {
	return t.ValidateEnum()
}

func (t targetVariableName) ValidateEnum() error {
	values := t.GetEnumValues()
	for _, v := range values {
		if v == string(t) {
			return nil
		}
	}
	return errors.New("")
}

func (t targetVariableName) GetEnumValues() []string {
	return targetVariableNames
}
