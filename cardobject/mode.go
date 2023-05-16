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

//Chosen Action
const ChosenAction = "CHOSEN ACTION"

//Chosen Entity
const ChosenEntity = "CHOSEN ENTITY"

//Chosen Place
const ChosenPlace = "CHOSEN PLACE"

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
var cardOppModes []string = []string{All, Random, Target}
var actionModes []string = []string{All, This, Random, Target, ChosenAction}
var entityModes []string = []string{All, This, Random, Target, ChosenEntity}
var placeModes []string = []string{All, This, Random, Target, ChosenPlace}
var actionOppModes []string = []string{All, Random, Target, ChosenAction}
var entityOppModes []string = []string{All, Random, Target, ChosenEntity}
var placeOppModes []string = []string{All, Random, Target, ChosenPlace}
var ownerModes []string = []string{Your, Opponents, Owners}

type IntChangeMode jsonschema.BasicEnum

func (i IntChangeMode) ValidateType(r jsonschema.RootElement) error {
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

func (s StringChangeMode) ValidateType(r jsonschema.RootElement) error {
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

func (p PlayerMode) ValidateType(r jsonschema.RootElement) error {
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

type CardMode jsonschema.TargetMode

func (c CardMode) ValidateType(r jsonschema.RootElement) error {
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

type CardOppMode jsonschema.TargetMode

func (c CardOppMode) ValidateType(r jsonschema.RootElement) error {
	values := c.EnumValues()
	for _, v := range values {
		if v == string(c) {
			return nil
		}
	}
	return errors.New("CardOppModes must be one of: " + strings.Join(cardOppModes, ","))
}

func (c CardOppMode) EnumValues() []string {
	return cardOppModes
}

type ActionMode jsonschema.TargetMode

func (a ActionMode) ValidateType(r jsonschema.RootElement) error {
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionModes must be one of: " + strings.Join(actionModes, ","))
}

func (a ActionMode) EnumValues() []string {
	return actionModes
}

type EntityMode jsonschema.TargetMode

func (e EntityMode) ValidateType(r jsonschema.RootElement) error {
	values := e.EnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityModes must be one of: " + strings.Join(entityModes, ","))
}

func (e EntityMode) EnumValues() []string {
	return entityModes
}

type PlaceMode jsonschema.TargetMode

func (p PlaceMode) ValidateType(r jsonschema.RootElement) error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceModes must be one of: " + strings.Join(placeModes, ","))
}

func (p PlaceMode) EnumValues() []string {
	return placeModes
}

type ActionOppMode jsonschema.TargetMode

func (a ActionOppMode) ValidateType(r jsonschema.RootElement) error {
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionOppModes must be one of: " + strings.Join(actionOppModes, ","))
}

func (a ActionOppMode) EnumValues() []string {
	return actionOppModes
}

type EntityOppMode jsonschema.TargetMode

func (e EntityOppMode) ValidateType(r jsonschema.RootElement) error {
	values := e.EnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityOppModes must be one of: " + strings.Join(entityOppModes, ","))
}

func (e EntityOppMode) EnumValues() []string {
	return entityOppModes
}

type PlaceOppMode jsonschema.TargetMode

func (p PlaceOppMode) ValidateType(r jsonschema.RootElement) error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceOppModes must be one of: " + strings.Join(placeOppModes, ","))
}

func (p PlaceOppMode) EnumValues() []string {
	return placeOppModes
}

type OwnerMode jsonschema.BasicEnum

func (o OwnerMode) ValidateType(r jsonschema.RootElement) error {
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
