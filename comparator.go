package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

var intComparators []string = []string{"EQUAL", "GREATER", "LESSER"}
var stringComparators []string = []string{"EQUAL", "CONTAINS", "UNEQUAL", "CONTAINSNOT"}

type intComparator struct{ *jsonschema.BasicEnum }

func (i intComparator) GetEnumValues() []string {
	return intComparators
}

type stringComparator struct{ *jsonschema.BasicEnum }

func (s stringComparator) GetEnumValues() []string {
	return stringComparators
}
