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
	Cost    *cost
	Effects []effect
}

type triggeredAbility struct {
	Cause   *eventListener
	Cost    *cost
	Effects []effect
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

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (a *activatedAbility) validate() error {
	errorRange := []error{}
	if a.Cost == nil {
		errorRange = append(errorRange, errors.New("ActivatedAbility must have Cost"))
	} else {
		errorRange = append(errorRange, a.Cost.validate())
	}
	errorRange = append(errorRange, validateEffects(a.Effects))
	return combineErrors(errorRange)
}

func (a *triggeredAbility) validate() error {
	errorRange := []error{}
	if a.Cause == nil {
		errorRange = append(errorRange, errors.New("TriggeredAbility must have Cause"))
	} else {
		errorRange = append(errorRange, a.Cause.validate())
	}
	if a.Cost != nil {
		errorRange = append(errorRange, a.Cost.validate())
	}
	errorRange = append(errorRange, validateEffects(a.Effects))
	return combineErrors(errorRange)
}
