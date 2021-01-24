package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var intChangeModes []string = []string{"INCREASES", "DECREASES", "CHANGES"}
var stringChangeModes []string = []string{"CHANGES"}

var playerModes []string = []string{"YOU", "OPPONENT"}
var cardModes []string = []string{"ALL", "OPPONENTSCHOICE", "RANDOM", "TARGET"}
var ownerModes []string = []string{"YOUR", "OPPONENTS", "OWNERS"}

type intChangeMode jsonschema.BasicEnum

func (i intChangeMode) Validate() error {
	return i.ValidateEnum()
}

func (i intChangeMode) ValidateEnum() error {
	values := i.GetEnumValues()
	for _, v := range values {
		if v == string(i) {
			return nil
		}
	}
	return errors.New("")
}

func (i intChangeMode) GetEnumValues() []string {
	return intChangeModes
}

type stringChangeMode jsonschema.BasicEnum

func (s stringChangeMode) Validate() error {
	return s.ValidateEnum()
}

func (s stringChangeMode) ValidateEnum() error {
	values := s.GetEnumValues()
	for _, v := range values {
		if v == string(s) {
			return nil
		}
	}
	return errors.New("")
}

func (s stringChangeMode) GetEnumValues() []string {
	return stringChangeModes
}

type playerMode jsonschema.BasicEnum

func (p playerMode) Validate() error {
	return p.ValidateEnum()
}

func (p playerMode) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p playerMode) GetEnumValues() []string {
	return playerModes
}

type cardMode jsonschema.BasicEnum

func (c cardMode) Validate() error {
	return c.ValidateEnum()
}

func (c cardMode) ValidateEnum() error {
	values := c.GetEnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("")
}

func (c cardMode) GetEnumValues() []string {
	return cardModes
}

type ownerMode jsonschema.BasicEnum

func (o ownerMode) Validate() error {
	return o.ValidateEnum()
}

func (o ownerMode) ValidateEnum() error {
	values := o.GetEnumValues()
	for _, v := range values {
		if v == string(o) {
			return nil
		}
	}
	return errors.New("")
}

func (o ownerMode) GetEnumValues() []string {
	return ownerModes
}
