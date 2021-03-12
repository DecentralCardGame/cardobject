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

type ActionIntProperty jsonschema.BasicEnum

func (a ActionIntProperty) Validate() error {
	return a.ValidateEnum()
}

func (a ActionIntProperty) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a ActionIntProperty) GetEnumValues() []string {
	return actionIntProperties
}

type ActionStringProperty jsonschema.BasicEnum

func (a ActionStringProperty) Validate() error {
	return a.ValidateEnum()
}

func (a ActionStringProperty) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a ActionStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type EntityIntProperty jsonschema.BasicEnum

func (e EntityIntProperty) Validate() error {
	return e.ValidateEnum()
}

func (e EntityIntProperty) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("")
}

func (e EntityIntProperty) GetEnumValues() []string {
	return entityIntProperties
}

type EntityStringProperty jsonschema.BasicEnum

func (e EntityStringProperty) Validate() error {
	return e.ValidateEnum()
}

func (e EntityStringProperty) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("")
}

func (e EntityStringProperty) GetEnumValues() []string {
	return entityStringProperties
}

type PlaceIntProperty jsonschema.BasicEnum

func (p PlaceIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlaceIntProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p PlaceIntProperty) GetEnumValues() []string {
	return placeIntProperties
}

type PlaceStringProperty jsonschema.BasicEnum

func (p PlaceStringProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlaceStringProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p PlaceStringProperty) GetEnumValues() []string {
	return actionStringProperties
}

type CardIntProperty jsonschema.BasicEnum

func (c CardIntProperty) Validate() error {
	return c.ValidateEnum()
}

func (c CardIntProperty) ValidateEnum() error {
	values := c.GetEnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("")
}

func (c CardIntProperty) GetEnumValues() []string {
	return cardIntProperties
}

type CardStringProperty jsonschema.BasicEnum

func (c CardStringProperty) Validate() error {
	return c.ValidateEnum()
}

func (c CardStringProperty) ValidateEnum() error {
	values := c.GetEnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("")
}

func (c CardStringProperty) GetEnumValues() []string {
	return cardStringProperties
}

type PlayerIntProperty jsonschema.BasicEnum

func (p PlayerIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlayerIntProperty) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p PlayerIntProperty) GetEnumValues() []string {
	return playerIntProperties
}
