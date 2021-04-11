package cardobject

import (
	"errors"
	"strings"

	"cardobject/jsonschema"
)

var intChangeModes []string = []string{"INCREASES", "DECREASES", "CHANGES"}
var stringChangeModes []string = []string{"CHANGES"}

var playerModes []string = []string{"YOU", "OPPONENT"}
var cardModes []string = []string{"ALL", "OPPONENTSCHOICE", "RANDOM", "TARGET"}
var ownerModes []string = []string{"YOUR", "OPPONENTS", "OWNERS"}

type IntChangeMode jsonschema.BasicEnum

func (i IntChangeMode) Validate() error {
	return i.ValidateEnum()
}

func (i IntChangeMode) ValidateEnum() error {
	values := i.EnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("IntChangeMode must be one of: " + strings.Join(intChangeModes, ","))
}

func (i IntChangeMode) EnumValues() []string {
	return intChangeModes
}

type StringChangeMode jsonschema.BasicEnum

func (s StringChangeMode) Validate() error {
	return s.ValidateEnum()
}

func (s StringChangeMode) ValidateEnum() error {
	values := s.EnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("StringComparator must be one of: " + strings.Join(stringComparators, ","))
}

func (s StringChangeMode) EnumValues() []string {
	return stringChangeModes
}

type PlayerMode jsonschema.BasicEnum

func (p PlayerMode) Validate() error {
	return p.ValidateEnum()
}

func (p PlayerMode) ValidateEnum() error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlayerModes must be one of: " + strings.Join(playerModes, ","))
}

func (p PlayerMode) EnumValues() []string {
	return playerModes
}

type CardMode jsonschema.BasicEnum

func (c CardMode) Validate() error {
	return c.ValidateEnum()
}

func (c CardMode) ValidateEnum() error {
	values := c.EnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("CardModes must be one of: " + strings.Join(cardModes, ","))
}

func (c CardMode) EnumValues() []string {
	return cardModes
}

type OwnerMode jsonschema.BasicEnum

func (o OwnerMode) Validate() error {
	return o.ValidateEnum()
}

func (o OwnerMode) ValidateEnum() error {
	values := o.EnumValues()
	for _, v := range values {
		if v == string(o) {
			return nil
		}
	}
	return errors.New("OwnerModes must be one of: " + strings.Join(ownerModes, ","))
}

func (o OwnerMode) EnumValues() []string {
	return ownerModes
}
