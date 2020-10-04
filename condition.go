package cardobject

import "errors"

type cardCondition struct {
	ActionCondition *actionCondition `json:",omitempty"`
	EntityCondition *entityCondition `json:",omitempty"`
	PlaceCondition  *placeCondition  `json:",omitempty"`
	ThisCondition   *thisCondition   `json:",omitempty"`
}

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

func (c *cardCondition) validate() error {
	possibleImplementer := []validateable{c.ActionCondition, c.EntityCondition, c.PlaceCondition, c.ThisCondition}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *actionCondition) validate() error {
	possibleImplementer := []validateable{c.ActionIntCondition, c.ActionStringCondition, c.ActionTagCondition}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *entityCondition) validate() error {
	possibleImplementer := []validateable{c.EntityIntCondition, c.EntityStringCondition, c.EntityTagCondition}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *placeCondition) validate() error {
	possibleImplementer := []validateable{c.PlaceIntCondition, c.PlaceStringCondition, c.PlaceTagCondition}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
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
