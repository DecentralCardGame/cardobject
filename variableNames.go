package cardobject

import "github.com/DecentralCardGame/jsonschema"

var intVariableNames []string = []string{"X", "Y", "Z"}
var stringVariableNames []string = []string{"A", "B", "C"}
var targetVariableNames []string = []string{"M", "T"}

type intVariableName struct{ *jsonschema.BasicEnum }

func (i intVariableName) GetEnumValues() []string {
	return intVariableNames
}

type stringVariableName struct{ *jsonschema.BasicEnum }

func (s stringVariableName) GetEnumValues() []string {
	return stringVariableNames
}

type targetVariableName struct{ *jsonschema.BasicEnum }

func (t targetVariableName) GetEnumValues() []string {
	return targetVariableNames
}
