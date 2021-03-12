package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var intComparators []string = []string{"EQUAL", "GREATER", "LESSER"}
var stringComparators []string = []string{"EQUAL", "CONTAINS", "UNEQUAL", "CONTAINSNOT"}

type IntComparator jsonschema.BasicEnum

func (i IntComparator) Validate() error {
	return i.ValidateEnum()
}

func (i IntComparator) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i IntComparator) GetEnumValues() []string {
	return intComparators
}

type StringComparator jsonschema.BasicEnum

func (s StringComparator) Validate() error {
	return s.ValidateEnum()
}

func (s StringComparator) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s StringComparator) GetEnumValues() []string {
	return stringComparators
}
