package cardobject

import (
	"errors"
	"strings"

	"cardobject/jsonschema"
)

var actionIntProperties []string = []string{"COSTSUM"}
var actionStringProperties []string = []string{"NAME", "TEXT"}
var entityIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH", "BASEHEALTH"}
var entityStringProperties []string = []string{"NAME", "TEXT"}
var headquarterIntProperties []string = []string{"HEALTH", "BASEHEALTH"}
var headquarterStringProperties []string = []string{"NAME", "TEXT"}
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
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionIntProperty must be one of: " + strings.Join(actionIntProperties, ","))
}

func (a ActionIntProperty) EnumValues() []string {
	return actionIntProperties
}

type ActionStringProperty jsonschema.BasicEnum

func (a ActionStringProperty) Validate() error {
	return a.ValidateEnum()
}

func (a ActionStringProperty) ValidateEnum() error {
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionStringProperty must be one of: " + strings.Join(actionStringProperties, ","))
}

func (a ActionStringProperty) EnumValues() []string {
	return actionStringProperties
}

type EntityIntProperty jsonschema.BasicEnum

func (e EntityIntProperty) Validate() error {
	return e.ValidateEnum()
}

func (e EntityIntProperty) ValidateEnum() error {
	values := e.EnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityIntProperty must be one of: " + strings.Join(entityIntProperties, ","))
}

func (e EntityIntProperty) EnumValues() []string {
	return entityIntProperties
}

type EntityStringProperty jsonschema.BasicEnum

func (e EntityStringProperty) Validate() error {
	return e.ValidateEnum()
}

func (e EntityStringProperty) ValidateEnum() error {
	values := e.EnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityStringProperty must be one of: " + strings.Join(entityStringProperties, ","))
}

func (e EntityStringProperty) EnumValues() []string {
	return entityStringProperties
}

type HeadquarterIntProperty jsonschema.BasicEnum

func (h HeadquarterIntProperty) Validate() error {
	return h.ValidateEnum()
}

func (h HeadquarterIntProperty) ValidateEnum() error {
	values := h.EnumValues()
	for _, v := range values {
		if v == string(h) {
			return nil
		}
	}
	return errors.New("HeadquarterIntProperty must be one of: " + strings.Join(headquarterIntProperties, ","))
}

func (h HeadquarterIntProperty) EnumValues() []string {
	return headquarterIntProperties
}

type HeadquarterStringProperty jsonschema.BasicEnum

func (h HeadquarterStringProperty) Validate() error {
	return h.ValidateEnum()
}

func (h HeadquarterStringProperty) ValidateEnum() error {
	values := h.EnumValues()
	for _, v := range values {
		if v == string(h) {
			return nil
		}
	}
	return errors.New("HeadquarterStringProperty must be one of: " + strings.Join(headquarterStringProperties, ","))
}

func (h HeadquarterStringProperty) EnumValues() []string {
	return headquarterStringProperties
}

type PlaceIntProperty jsonschema.BasicEnum

func (p PlaceIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlaceIntProperty) ValidateEnum() error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceIntProperty must be one of: " + strings.Join(placeIntProperties, ","))
}

func (p PlaceIntProperty) EnumValues() []string {
	return placeIntProperties
}

type PlaceStringProperty jsonschema.BasicEnum

func (p PlaceStringProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlaceStringProperty) ValidateEnum() error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceStringProperty must be one of: " + strings.Join(placeStringProperties, ","))
}

func (p PlaceStringProperty) EnumValues() []string {
	return placeStringProperties
}

type CardIntProperty jsonschema.BasicEnum

func (c CardIntProperty) Validate() error {
	return c.ValidateEnum()
}

func (c CardIntProperty) ValidateEnum() error {
	values := c.EnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("CardIntProperty must be one of: " + strings.Join(cardIntProperties, ","))
}

func (c CardIntProperty) EnumValues() []string {
	return cardIntProperties
}

type CardStringProperty jsonschema.BasicEnum

func (c CardStringProperty) Validate() error {
	return c.ValidateEnum()
}

func (c CardStringProperty) ValidateEnum() error {
	values := c.EnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("CardStringProperty must be one of: " + strings.Join(cardStringProperties, ","))
}

func (c CardStringProperty) EnumValues() []string {
	return cardStringProperties
}

type PlayerIntProperty jsonschema.BasicEnum

func (p PlayerIntProperty) Validate() error {
	return p.ValidateEnum()
}

func (p PlayerIntProperty) ValidateEnum() error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlayerIntProperty must be one of: " + strings.Join(playerIntProperties, ","))
}

func (p PlayerIntProperty) EnumValues() []string {
	return playerIntProperties
}
