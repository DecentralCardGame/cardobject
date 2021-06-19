package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//BaseAttack Property
const BaseAttack = "BASEATTACK"

//BaseCostSum Property
const BaseCostSum = "BASECOSTSUM"

//BaseHealth Property
const BaseHealth = "BASEHEALTH"

//CurrentAttack Property
const CurrentAttack = "ATTACK"

//CurrentCostSum Property
const CurrentCostSum = "COSTSUM"

//CurrentHealth Property
const CurrentHealth = "HEALTH"

//Name Property
const Name = "NAME"

//Text Property
const Text = "TEXT"

var actionIntProperties []string = []string{BaseCostSum, CurrentCostSum}
var actionStringProperties []string = []string{Name, Text}
var entityIntProperties []string = []string{CurrentAttack, BaseAttack, BaseCostSum, BaseHealth, CurrentCostSum, CurrentHealth}
var entityStringProperties []string = []string{Name, Text}
var headquarterIntProperties []string = []string{BaseHealth, CurrentHealth}
var headquarterStringProperties []string = []string{Name, Text}
var placeIntProperties []string = []string{BaseCostSum, BaseHealth, CurrentCostSum, CurrentHealth}
var placeStringProperties []string = []string{Name, Text}

var cardIntProperties []string = []string{"ATTACK", "COSTSUM", "HEALTH"}
var cardStringProperties []string = []string{"NAME", "TEXT"}

var playerIntProperties []string = []string{"HANDSIZE", "BOARDSIZE", "DECKSIZE"}

type ActionIntProperty jsonschema.BasicEnum

func (a ActionIntProperty) ValidateType(r jsonschema.RootElement) error {
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

func (a ActionStringProperty) ValidateType(r jsonschema.RootElement) error {
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

func (e EntityIntProperty) ValidateType(r jsonschema.RootElement) error {
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

func (e EntityStringProperty) ValidateType(r jsonschema.RootElement) error {
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

func (h HeadquarterIntProperty) ValidateType(r jsonschema.RootElement) error {
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

func (h HeadquarterStringProperty) ValidateType(r jsonschema.RootElement) error {
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

func (p PlaceIntProperty) ValidateType(r jsonschema.RootElement) error {
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

func (p PlaceStringProperty) ValidateType(r jsonschema.RootElement) error {
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

func (c CardIntProperty) ValidateType(r jsonschema.RootElement) error {
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

func (c CardStringProperty) ValidateType(r jsonschema.RootElement) error {
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

func (p PlayerIntProperty) ValidateType(r jsonschema.RootElement) error {
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
