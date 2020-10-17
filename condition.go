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

type playerConditions []playerCondition

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

type playerCondition struct {
	PlayerIntCondition *playerIntCondition `json:",omitempty"`
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

type playerIntCondition struct {
	PlayerIntProperty string
	IntValue          int
	IntComparator     string
}

func (c *cardConditions) validate() error {
	possibleImplementer := []validateable{c.ActionConditions, c.EntityConditions, c.PlaceConditions, c.ThisConditions}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Conditions implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c actionConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " actionConditions")
	}
	errorRange := []error{}
	for _, actionCondition := range c {
		errorRange = append(errorRange, actionCondition.validate())
	}
	return combineErrors(errorRange)
}

func (c entityConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " entityConditions")
	}
	errorRange := []error{}
	for _, entityCondition := range c {
		errorRange = append(errorRange, entityCondition.validate())
	}
	return combineErrors(errorRange)
}

func (c placeConditions) validate() error {
	if len(c) > maxConditionCount {
		return errors.New("The card must have at most " + strconv.Itoa(maxConditionCount) + " placeConditions")
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
		return errors.New("ActionCondition implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *entityCondition) validate() error {
	possibleImplementer := []validateable{c.EntityIntCondition, c.EntityStringCondition, c.EntityTagCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("EntityCondition implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *placeCondition) validate() error {
	possibleImplementer := []validateable{c.PlaceIntCondition, c.PlaceStringCondition, c.PlaceTagCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("PlaceCondition implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *playerCondition) validate() error {
	possibleImplementer := []validateable{c.PlayerIntCondition}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("PlayerCondition implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *thisCondition) validate() error {
	return nil
}

func (c *actionIntCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionIntProperty(c.ActionIntProperty))
	errorRange = append(errorRange, validateIntComparator(c.IntComparator))
	errorRange = append(errorRange, validateSimpleInt(c.IntValue))
	return combineErrors(errorRange)
}

func (c *actionStringCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateActionStringProperty(c.ActionStringProperty))
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *actionTagCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *entityIntCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityIntProperty(c.EntityIntProperty))
	errorRange = append(errorRange, validateIntComparator(c.IntComparator))
	errorRange = append(errorRange, validateSimpleInt(c.IntValue))
	return combineErrors(errorRange)
}

func (c *entityStringCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateEntityStringProperty(c.EntityStringProperty))
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *entityTagCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *placeIntCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceIntProperty(c.PlaceIntProperty))
	errorRange = append(errorRange, validateIntComparator(c.IntComparator))
	errorRange = append(errorRange, validateSimpleInt(c.IntValue))
	return combineErrors(errorRange)
}

func (c *placeStringCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlaceStringProperty(c.PlaceStringProperty))
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *placeTagCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateStringComparator(c.StringComparator))
	errorRange = append(errorRange, validateSimpleString(c.StringValue))
	return combineErrors(errorRange)
}

func (c *playerIntCondition) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validatePlayerIntProperty(c.PlayerIntProperty))
	errorRange = append(errorRange, validateIntComparator(c.IntComparator))
	errorRange = append(errorRange, validateSimpleInt(c.IntValue))
	return combineErrors(errorRange)
}
