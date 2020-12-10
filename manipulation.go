package cardobject

import (
	"errors"
	"strconv"
)

type manipulationsInterface struct {
	ActionManipulations *actionManipulations `json:",omitempty"`
	EntityManipulations *entityManipulations `json:",omitempty"`
	PlaceManipulations  *placeManipulations  `json:",omitempty"`
}

type actionManipulations []actionManipulationInterface

type entityManipulations []entityManipulationInterface

type placeManipulations []placeManipulationInterface

type actionManipulationInterface struct {
	ActionEffectManipulation *actionEffectManipulation `json:",omitempty"`
	ActionIntManipulation    *actionIntManipulation    `json:",omitempty"`
	ActionStringManipulation *actionStringManipulation `json:",omitempty"`
	ActionTagManipulation    *actionTagManipulation    `json:",omitempty"`
	ActionZoneChange         *actionZoneChange         `json:",omitempty"`
}

type entityManipulationInterface struct {
	EntityAbilityManipulation *entityAbilityManipulation `json:",omitempty"`
	EntityIntManipulation     *entityIntManipulation     `json:",omitempty"`
	EntityStringManipulation  *entityStringManipulation  `json:",omitempty"`
	EntityTagManipulation     *entityTagManipulation     `json:",omitempty"`
	EntityZoneChange          *entityZoneChange          `json:",omitempty"`
}

type placeManipulationInterface struct {
	PlaceAbilityManipulation *placeAbilityManipulation `json:",omitempty"`
	PlaceIntManipulation     *placeIntManipulation     `json:",omitempty"`
	PlaceStringManipulation  *placeStringManipulation  `json:",omitempty"`
	PlaceTagManipulation     *placeTagManipulation     `json:",omitempty"`
	PlaceZoneChange          *placeZoneChange          `json:",omitempty"`
}

type actionEffectManipulation struct {
	Effect         effectInterface
	EffectOperator string
}

type actionIntManipulation struct {
	IntProperty string
	IntOperator string
	IntValue    intValueInterface
}

type actionStringManipulation struct {
	StringProperty string
	StringOperator string
	StringValue    string
}

type actionTagManipulation struct {
	StringValue    string
	StringOperator string
}

type entityAbilityManipulation struct {
	Ability         abilityInterface
	AbilityOperator string
}

type entityIntManipulation struct {
	IntProperty string
	IntOperator string
	IntValue    intValueInterface
}

type entityStringManipulation struct {
	StringProperty string
	StringOperator string
	StringValue    string
}

type entityTagManipulation struct {
	StringValue    string
	StringOperator string
}

type placeAbilityManipulation struct {
	Ability         abilityInterface
	AbilityOperator string
}

type placeIntManipulation struct {
	IntProperty string
	IntOperator string
	IntValue    intValueInterface
}

type placeStringManipulation struct {
	StringProperty string
	StringOperator string
	StringValue    string
}

type placeTagManipulation struct {
	StringValue    string
	StringOperator string
}

func (m *manipulationsInterface) validate() error {
	possibleImplementer := []validateable{m.ActionManipulations, m.EntityManipulations, m.PlaceManipulations}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Manipulations implemented by not exactly one option")
	}
	return implementer.validate()
}

func (m actionManipulations) validate() error {
	if len(m) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxManipulationCount) + " ActionManipulations")
	}
	errorRange := []error{}
	for _, actionManipulation := range m {
		errorRange = append(errorRange, actionManipulation.validate())
	}
	return combineErrors(errorRange)
}

func (m entityManipulations) validate() error {
	if len(m) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxManipulationCount) + " EntityManipulations")
	}
	errorRange := []error{}
	for _, entityManipulation := range m {
		errorRange = append(errorRange, entityManipulation.validate())
	}
	return combineErrors(errorRange)
}

func (m placeManipulations) validate() error {
	if len(m) > maxExtractorCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxManipulationCount) + " EntityManipulations")
	}
	errorRange := []error{}
	for _, placeManipulation := range m {
		errorRange = append(errorRange, placeManipulation.validate())
	}
	return combineErrors(errorRange)
}

func (m *actionManipulationInterface) validate() error {
	possibleImplementer := []validateable{m.ActionEffectManipulation, m.ActionIntManipulation, m.ActionStringManipulation, m.ActionTagManipulation, m.ActionZoneChange}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("ActionManipulation implemented by not exactly one option")
	}
	return implementer.validate()
}

func (m *entityManipulationInterface) validate() error {
	possibleImplementer := []validateable{m.EntityAbilityManipulation, m.EntityIntManipulation, m.EntityStringManipulation, m.EntityTagManipulation, m.EntityZoneChange}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("EntityManipulation implemented by not exactly one option")
	}
	return implementer.validate()
}

func (m *placeManipulationInterface) validate() error {
	possibleImplementer := []validateable{m.PlaceAbilityManipulation, m.PlaceIntManipulation, m.PlaceStringManipulation, m.PlaceTagManipulation, m.PlaceZoneChange}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("PlaceManipulation implemented by not exactly one option")
	}
	return implementer.validate()
}

func (m *actionEffectManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, m.Effect.validate())
	errorRange = append(errorRange, validateAbilityEffectOperator(m.EffectOperator))
	return combineErrors(errorRange)
}

func (m *actionIntManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionIntProperty(m.IntProperty))
	errorRange = append(errorRange, validateIntOperator(m.IntOperator))
	errorRange = append(errorRange, m.IntValue.validate())
	return combineErrors(errorRange)
}

func (m *actionStringManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionStringProperty(m.StringProperty))
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}

func (m *actionTagManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}

func (m *entityAbilityManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, m.Ability.validate())
	errorRange = append(errorRange, validateAbilityEffectOperator(m.AbilityOperator))
	return combineErrors(errorRange)
}

func (m *entityIntManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityIntProperty(m.IntProperty))
	errorRange = append(errorRange, validateIntOperator(m.IntOperator))
	errorRange = append(errorRange, m.IntValue.validate())
	return combineErrors(errorRange)
}

func (m *entityStringManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityStringProperty(m.StringProperty))
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}

func (m *entityTagManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}

func (m *placeAbilityManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, m.Ability.validate())
	errorRange = append(errorRange, validateAbilityEffectOperator(m.AbilityOperator))
	return combineErrors(errorRange)
}

func (m *placeIntManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceIntProperty(m.IntProperty))
	errorRange = append(errorRange, validateIntOperator(m.IntOperator))
	errorRange = append(errorRange, m.IntValue.validate())
	return combineErrors(errorRange)
}

func (m *placeStringManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceStringProperty(m.StringProperty))
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}

func (m *placeTagManipulation) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringOperator(m.StringOperator))
	errorRange = append(errorRange, validateSimpleString(m.StringValue))
	return combineErrors(errorRange)
}
