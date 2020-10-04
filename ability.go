package cardobject

import (
	"errors"
	"strconv"
)

type ability struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}

type activatedAbility struct {
	Cost    *cost `json:",omitempty"`
	Effects []effect
}

type triggeredAbility struct {
}

func validateAbilities(abilities []ability) error {
	if len(abilities) > maxAbilityEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxAbilityEffectCount) + " effects")
	}
	errorRange := []error{}
	for _, ability := range abilities {
		errorRange = append(errorRange, ability.validate())
	}
	return combineErrors(errorRange)
}

func (a *ability) validate() error {
	possibleImplementer := []validateable{a.ActivatedAbility, a.TriggeredAbility}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (a *activatedAbility) validate() error {
	errors := []error{}
	errors = append(errors, a.Cost.validate())
	errors = append(errors, validateEffects(a.Effects))
	return combineErrors(errors)
}

func (a *triggeredAbility) validate() error {
	return nil
}
