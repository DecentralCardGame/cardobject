package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//All Mode
const All = "ALL"

//Changes Mode
const Changes = "CHANGES"

//Decreases Mode
const Decreases = "DECREASES"

//Increases Mode
const Increases = "INCREASES"

//Opponent Mode
const Opponent = "OPPONENT"

//Opponents Mode
const Opponents = "OPPONENTS"

//OpponentsChoice Mode
const OpponentsChoice = "OPPONENTSCHOICE"

//Owners Mode
const Owners = "OWNERS"

//Random Mode
const Random = "RANDOM"

//Target Mode
const Target = "TARGET"

//This Mode
const This = "THIS"

//You Mode
const You = "YOU"

//Your Mode
const Your = "YOUR"

var intChangeModes []string = []string{Increases, Decreases, Changes}
var stringChangeModes []string = []string{Changes}

var playerModes []string = []string{You, Opponent}
var cardModes []string = []string{All, This, Random, Target}
var ownerModes []string = []string{Your, Opponents, Owners}

type IntChangeMode jsonschema.BasicEnum

func (i IntChangeMode) Validate(r jsonschema.RootElement) error {
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

func (s StringChangeMode) Validate(r jsonschema.RootElement) error {
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

func (p PlayerMode) Validate(r jsonschema.RootElement) error {
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

func (c CardMode) Validate(r jsonschema.RootElement) error {
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

func (o OwnerMode) Validate(r jsonschema.RootElement) error {
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
