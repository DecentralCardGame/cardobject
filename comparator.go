package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/jsonschema"
)

var intComparators []string = []string{"EQUAL", "GREATER", "LESSER"}
var stringComparators []string = []string{"EQUAL", "CONTAINS", "UNEQUAL", "CONTAINSNOT"}

type IntComparator jsonschema.BasicEnum

func (i IntComparator) Validate() error {
	return i.ValidateEnum()
}

func (i IntComparator) ValidateEnum() error {
	values := i.EnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("IntComparator must be one of: " + strings.Join(intComparators, ","))
}

func (i IntComparator) EnumValues() []string {
	return intComparators
}

type StringComparator jsonschema.BasicEnum

func (s StringComparator) Validate() error {
	return s.ValidateEnum()
}

func (s StringComparator) ValidateEnum() error {
	values := s.EnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("StringComparator must be one of: " + strings.Join(stringComparators, ","))
}

func (s StringComparator) EnumValues() []string {
	return stringComparators
}
