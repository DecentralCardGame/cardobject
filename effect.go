package cardobject

import (
	"errors"
	"strconv"
)

type effect struct {
}

func validateEffects(effects []effect) error {
	if len(effects) > maxAbilityEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxAbilityEffectCount) + " effects")
	}
	errorRange := []error{}
	for _, effect := range effects {
		errorRange = append(errorRange, effect.validate())
	}
	return combineErrors(errorRange)
}

func (effect *effect) validate() error {
	return nil
}
