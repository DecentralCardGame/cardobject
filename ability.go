package cardobject

import (
	"errors"
	"strconv"
)

type abilities []abilityInterface

type abilityInterface struct {
	ActivatedAbility *activatedAbility `json:",omitempty"`
	TriggeredAbility *triggeredAbility `json:",omitempty"`
}

type activatedAbility struct {
	AbilityCost *costInterface
	Effects     effects
}

type triggeredAbility struct {
	Cause   *eventListenerInterface
	Cost    *costInterface
	Effects effects
}

func (a abilities) validate() error {
	if len(a) > maxAbilityEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxAbilityEffectCount) + " effects")
	}
	errorRange := []error{}
	for _, ability := range a {
		errorRange = append(errorRange, ability.validate())
	}
	return combineErrors(errorRange)
}

func (a *abilityInterface) validate() error {
	possibleImplementer := []validateable{a.ActivatedAbility, a.TriggeredAbility}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (a *activatedAbility) validate() error {
	errorRange := []error{}
	if a.AbilityCost == nil {
		errorRange = append(errorRange, errors.New("ActivatedAbility must have AbilityCost"))
	} else {
		errorRange = append(errorRange, a.AbilityCost.validate())
	}
	errorRange = append(errorRange, a.Effects.validate())
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
	errorRange = append(errorRange, a.Effects.validate())
	return combineErrors(errorRange)
}
