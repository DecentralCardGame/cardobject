package cardobject

import (
	"errors"
	"strconv"
)

type effects []effect

type effect struct {
	ProductionEffect *productionEffect `json:",omitempty"`
	DrawEffect       *drawEffect       `json:",omitempty"`
	TokenEffect      *tokenEffect      `json:",omitempty"`
	TargetEffect     *targetEffect     `json:",omitempty"`
	ChooseFromEffect *chooseFromEffect `json:",omitempty"`
}

type productionEffect struct {
	Amount intValue
}

type drawEffect struct {
	DrawAmount intValue
}

type tokenEffect struct {
	TokenAmount intValue
	Token       token
}

type chooseFromEffect struct {
	Effects effects
}

type targetEffect struct {
	ActionTargetEffect    *actionTargetEffect    `json:",omitempty"`
	EntityTargetEffect    *entityTargetEffect    `json:",omitempty"`
	PlaceTargetEffect     *placeTargetEffect     `json:",omitempty"`
	ExtractorTargetEffect *extractorTargetEffect `json:",omitempty"`
}

type actionTargetEffect struct {
	ActionSelector      *actionSelector
	ActionManipulations *actionManipulations
}

type entityTargetEffect struct {
	EntitySelector      *entitySelector
	EntityManipulations *entityManipulations
}

type placeTargetEffect struct {
	PlaceSelector      *placeSelector
	PlaceManipulations *placeManipulations
}

type extractorTargetEffect struct {
	TargetVariable string
	Manipulations  *manipulations
}

func (e effects) validate() error {
	if len(e) > maxAbilityEffectCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxAbilityEffectCount) + " effects")
	}
	errorRange := []error{}
	for _, effect := range e {
		errorRange = append(errorRange, effect.validate())
	}
	return combineErrors(errorRange)
}

func (e *effect) validate() error {
	possibleImplementer := []validateable{e.ChooseFromEffect, e.DrawEffect, e.ProductionEffect, e.TargetEffect, e.TokenEffect}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Effect implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *chooseFromEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.Effects.validate())
	return combineErrors(errorRange)
}

func (e *drawEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.DrawAmount.validate())
	return combineErrors(errorRange)
}

func (e *productionEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.Amount.validate())
	return combineErrors(errorRange)
}

func (e *tokenEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.TokenAmount.validate())
	errorRange = append(errorRange, e.Token.validate())
	return combineErrors(errorRange)
}

func (e *targetEffect) validate() error {
	possibleImplementer := []validateable{e.ActionTargetEffect, e.EntityTargetEffect, e.ExtractorTargetEffect, e.PlaceTargetEffect}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("TargetEffect implemented by not exactly one option")
	}
	return implementer.validate()
}

func (e *actionTargetEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.ActionSelector.validate())
	errorRange = append(errorRange, e.ActionManipulations.validate())
	return combineErrors(errorRange)
}

func (e *entityTargetEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.EntitySelector.validate())
	errorRange = append(errorRange, e.EntityManipulations.validate())
	return combineErrors(errorRange)
}

func (e *placeTargetEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, e.PlaceSelector.validate())
	errorRange = append(errorRange, e.PlaceManipulations.validate())
	return combineErrors(errorRange)
}

func (e *extractorTargetEffect) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateTargetVariableName(e.TargetVariable))
	errorRange = append(errorRange, e.Manipulations.validate())
	return combineErrors(errorRange)
}
