package blockchaincard

import (
	"errors"
	"strconv"
)

type effect struct {
}

func validateEffects(effects []effect) error {
	if len(effects) > maxEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxEffectCount) + " effects")
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
