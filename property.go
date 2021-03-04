package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var actionIntProperties []string = []string{"COSTSUM"}
var actionStringProperties []string = []string{"NAME", "TEXT"}
var entityIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH", "BASEHEALTH"}
var entityStringProperties []string = []string{"NAME", "TEXT"}
var placeIntProperties []string = []string{"COSTSUM", "HEALTH", "BASEHEALTH"}
var placeStringProperties []string = []string{"NAME", "TEXT"}

var cardIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var cardStringProperties []string = []string{"NAME", "TEXT"}

var playerIntProperties []string = []string{"HANDSIZE", "BOARDSIZE", "DECKSIZE"}

type actionIntProperty jsonschema.BasicEnum

func (a actionIntProperty) Validate() error {
	return a.ValidateEnum()
}

func (a actionIntProperty) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a actionIntProperty) GetEnumValues() []string {
	return actionIntProperties
}

type actionStringProperty jsonschema.BasicEnum

func (a actionStringProperty) Validate() error {
	return a.ValidateEnum()
}

func (a actionStringProperty) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a actionStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type entityIntProperty jsonschema.BasicEnum

func (e entityIntProperty) Validate() error {
	return e.ValidateEnum()
}

func (e entityIntProperty) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("")
}

func (e entityIntProperty) GetEnumValues() []string {
	return entityIntProperties
}

type entityStringProperty jsonschema.BasicEnum

func (e entityStringProperty) Validate() error {
	return e.ValidateEnum()
}

func (e entityStringProperty) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("")
}

func (e entityStringProperty) GetEnumValues() []string {
	return entityStringProperties
}

type placeIntProperty jsonschema.BasicEnum

func (p placeIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p placeIntProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p placeIntProperty) GetEnumValues() []string {
	return placeIntProperties
}

type placeStringProperty jsonschema.BasicEnum

func (p placeStringProperty) Validate() error {
	return p.ValidateEnum()
}

func (p placeStringProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p placeStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type cardIntProperty jsonschema.BasicEnum

func (c cardIntProperty) Validate() error {
	return c.ValidateEnum()
}

func (c cardIntProperty) ValidateEnum() error {
	values := c.GetEnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("")
}

func (c cardIntProperty) GetEnumValues() []string {
	return cardIntProperties
}

type cardStringProperty jsonschema.BasicEnum

func (c cardStringProperty) Validate() error {
	return c.ValidateEnum()
}

func (c cardStringProperty) ValidateEnum() error {
	values := c.GetEnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("")
}

func (c cardStringProperty) GetEnumValues() []string {
	return cardStringProperties
}

type playerIntProperty jsonschema.BasicEnum

func (p playerIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p playerIntProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p playerIntProperty) GetEnumValues() []string {
	return playerIntProperties
}
