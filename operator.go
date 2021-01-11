package cardobject

import "github.com/DecentralCardGame/jsonschema"

var abilityEffectOperators []string = []string{"GAIN", "LOSE"}
var intOperators []string = []string{"SET", "ADD", "SUBTRACT"}
var stringOperators []string = []string{"SET"}

var arithOperators []string = []string{"ADD", "SUBTRACT"}

type abilityEffectOperator struct{ *jsonschema.BasicEnum }

func (a abilityEffectOperator) GetEnumValues() []string {
	return abilityEffectOperators
}

type intOperator struct{ *jsonschema.BasicEnum }

func (i intOperator) GetEnumValues() []string {
	return intOperators
}

type stringOperator struct{ *jsonschema.BasicEnum }

func (s stringOperator) GetEnumValues() []string {
	return stringOperators
}

type arithOperator struct{ *jsonschema.BasicEnum }

func (a arithOperator) GetEnumValues() []string {
	return arithOperators
}
