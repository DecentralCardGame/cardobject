package cardobject

import "errors"

var abilityEffectOperators []string = []string{"GAIN", "LOSE"}
var intOperators []string = []string{"SET", "ADD", "SUBTRACT"}
var stringOperators []string = []string{"SET"}

var arithOperators []string = []string{"ADD", "SUBTRACT"}

func validateAbilityEffectOperator(property string) error {
	for _, t := range abilityEffectOperators {
		if t == property {
			return nil
		}
	}
	return errors.New("AbilityEffectOperator: " + property + " not available")
}

func validateIntOperator(property string) error {
	for _, t := range intOperators {
		if t == property {
			return nil
		}
	}
	return errors.New("IntOperator: " + property + " not available")
}

func validateStringOperator(property string) error {
	for _, t := range stringOperators {
		if t == property {
			return nil
		}
	}
	return errors.New("StringOperator: " + property + " not available")
}

func validateArithOperator(property string) error {
	for _, t := range arithOperators {
		if t == property {
			return nil
		}
	}
	return errors.New("ArithOperator: " + property + " not available")
}
