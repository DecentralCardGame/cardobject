package cardobject

import (
	"errors"
	"strconv"
)

type cardConditions struct {
	ActionConditions *actionConditions `json:",omitempty"`
	EntityConditions *entityConditions `json:",omitempty"`
	PlaceConditions  *placeConditions  `json:",omitempty"`
	ThisConditions   *thisCondition    `json:",omitempty"`
}

type actionConditions []actionCondition

type entityConditions []entityCondition

type placeConditions []placeCondition

type actionCondition struct {
	ActionIntCondition    *actionIntCondition    `json:",omitempty"`
	ActionStringCondition *actionStringCondition `json:",omitempty"`
	ActionTagCondition    *actionTagCondition    `json:",omitempty"`
}

type entityCondition struct {
	EntityIntCondition    *entityIntCondition    `json:",omitempty"`
	EntityStringCondition *entityStringCondition `json:",omitempty"`
	EntityTagCondition    *entityTagCondition    `json:",omitempty"`
}

type placeCondition struct {
	PlaceIntCondition    *placeIntCondition    `json:",omitempty"`
	PlaceStringCondition *placeStringCondition `json:",omitempty"`
	PlaceTagCondition    *placeTagCondition    `json:",omitempty"`
}

type thisCondition struct{}

type actionIntCondition struct {
	ActionIntProperty string
	IntValue          int
	IntComparator     string
}

type actionStringCondition struct {
	ActionStringProperty string
	StringValue          string
	StringComparator     string
}

type actionTagCondition struct {
	StringValue      string
	StringComparator string
}

type entityIntCondition struct {
	EntityIntProperty string
	IntValue          int
	IntComparator     string
}

type entityStringCondition struct {
	EntityStringProperty string
	StringValue          string
	StringComparator     string
}

type entityTagCondition struct {
	StringValue      string
	StringComparator string
}

type placeIntCondition struct {
	PlaceIntProperty string
	IntValue         int
	IntComparator    string
}

type placeStringCondition struct {
	PlaceStringProperty string
	StringValue         string
	StringComparator    string
}

type placeTagCondition struct {
	StringValue      string
	StringComparator string
}

func (c *cardConditions) validate() error {
	possibleImplementer := []validateable{c.ActionConditions, c.EntityConditions, c.PlaceConditions, c.ThisConditions}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c actionConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " conditions")
	}
	errorRange := []error{}
	for _, actionCondition := range c {
		errorRange = append(errorRange, actionCondition.validate())
	}
	return combineErrors(errorRange)
}

func (c entityConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " conditions")
	}
	errorRange := []error{}
	for _, entityCondition := range c {
		errorRange = append(errorRange, entityCondition.validate())
	}
	return combineErrors(errorRange)
}

func (c placeConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " conditions")
	}
	errorRange := []error{}
	for _, placeCondition := range c {
		errorRange = append(errorRange, placeCondition.validate())
	}
	return combineErrors(errorRange)
}

func (c *actionCondition) validate() error {
	possibleImplementer := []validateable{c.ActionIntCondition, c.ActionStringCondition, c.ActionTagCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *entityCondition) validate() error {
	possibleImplementer := []validateable{c.EntityIntCondition, c.EntityStringCondition, c.EntityTagCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *placeCondition) validate() error {
	possibleImplementer := []validateable{c.PlaceIntCondition, c.PlaceStringCondition, c.PlaceTagCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *thisCondition) validate() error {
	return nil
}

func (c *actionIntCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateActionIntProperty(c.ActionIntProperty))
	errors = append(errors, validateIntComparator(c.IntComparator))
	errors = append(errors, validateSimpleInt(c.IntValue))
	return combineErrors(errors)
}

func (c *actionStringCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateActionStringProperty(c.ActionStringProperty))
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}

func (c *actionTagCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}

func (c *entityIntCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateEntityIntProperty(c.EntityIntProperty))
	errors = append(errors, validateIntComparator(c.IntComparator))
	errors = append(errors, validateSimpleInt(c.IntValue))
	return combineErrors(errors)
}

func (c *entityStringCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateEntityStringProperty(c.EntityStringProperty))
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}

func (c *entityTagCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}

func (c *placeIntCondition) validate() error {
	errors := []error{}
	errors = append(errors, validatePlaceIntProperty(c.PlaceIntProperty))
	errors = append(errors, validateIntComparator(c.IntComparator))
	errors = append(errors, validateSimpleInt(c.IntValue))
	return combineErrors(errors)
}

func (c *placeStringCondition) validate() error {
	errors := []error{}
	errors = append(errors, validatePlaceStringProperty(c.PlaceStringProperty))
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}

func (c *placeTagCondition) validate() error {
	errors := []error{}
	errors = append(errors, validateStringComparator(c.StringComparator))
	errors = append(errors, validateSimpleString(c.StringValue))
	return combineErrors(errors)
}
