package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var intComparators []string = []string{"EQUAL", "GREATER", "LESSER"}
var stringComparators []string = []string{"EQUAL", "CONTAINS", "UNEQUAL", "CONTAINSNOT"}

type intComparator jsonschema.BasicEnum

func (i intComparator) Validate() error {
	return i.ValidateEnum()
}

func (i intComparator) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i intComparator) GetEnumValues() []string {
	return intComparators
}

type stringComparator jsonschema.BasicEnum

func (s stringComparator) Validate() error {
	return s.ValidateEnum()
}

func (s stringComparator) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s stringComparator) GetEnumValues() []string {
	return stringComparators
}
