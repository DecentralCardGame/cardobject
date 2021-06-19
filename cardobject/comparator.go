package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//Contains Comparator
const Contains = "CONTAINS"

//ContainsNot Comparator
const ContainsNot = "CONTAINSNOT"

//Equal Comparator
const Equal = "EQUAL"

//Greater Comparator
const Greater = "GREATER"

//Lesser Comparator
const Lesser = "LESSER"

//Uneaqual Comparator
const Uneaqual = "UNEQUAL"

var intComparators []string = []string{Equal, Greater, Lesser}
var stringComparators []string = []string{Equal, Contains, Uneaqual, ContainsNot}

type IntComparator jsonschema.BasicEnum

func (i IntComparator) Validate(r jsonschema.RootElement) error {
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

func (s StringComparator) Validate(r jsonschema.RootElement) error {
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
