package blockchainCard

import (
	"errors"
	"strconv"
)

type ability struct {
}

func validateAbilities(abilities []ability) error {
	if len(abilities) > maxEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxEffectCount) + " effects")
	}
	errorRange := []error{}
	for _, ability := range abilities {
		errorRange = append(errorRange, ability.validate())
	}
	return combineErrors(errorRange)
}

func (ability *ability) validate() error {
	return nil
}
