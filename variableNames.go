package cardobject

import "errors"

var intVariableNames []string = []string{"X", "Y", "Z"}
var stringVariableNames []string = []string{"A", "B", "C"}
var targetVariableNames []string = []string{"M", "T"}

func validateIntVariableName(variableName string) error {
	for _, t := range intVariableNames {
		if t == variableName {
			return nil
		}
	}
	return errors.New("IntVariableName: " + variableName + " not available")
}

func validateStringVariableName(variableName string) error {
	for _, t := range stringVariableNames {
		if t == variableName {
			return nil
		}
	}
	return errors.New("StringVariableName: " + variableName + " not available")
}

func validateTargetVariableName(variableName string) error {
	for _, t := range targetVariableNames {
		if t == variableName {
			return nil
		}
	}
	return errors.New("TargetVariableName: " + variableName + " not available")
}
