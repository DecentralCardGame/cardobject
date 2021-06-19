package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//AttackLane Zone
const AttackLane = "ATTACKLANE"

//BlockLane Zone
const BlockLane = "BLOCKLANE"

//Deck Zone
const Deck = "DECK"

//Dustpile Zone
const Dustpile = "DUSTPILE"

//Field Zone
const Field = "FIELD"

//Hand Zone
const Hand = "HAND"

//Void Zone
const Void = "VOID"

var dynamicZones []string = []string{Deck, Dustpile, Field, Hand}
var zones []string = []string{Deck, Dustpile, Field, Hand, Void}

var actionZones []string = []string{Deck, Dustpile, Hand, Void}
var entityZones []string = []string{AttackLane, BlockLane, Deck, Dustpile, Field, Hand, Void}
var placeZones []string = []string{Deck, Dustpile, Field, Hand, Void}

type DynamicZone jsonschema.BasicEnum

func (d DynamicZone) Validate(r jsonschema.RootElement) error {
	values := d.EnumValues()
	for _, v := range values {
		if v == string(d) {
			return nil
		}
	}
	return errors.New("DynamicZone must be on of " + strings.Join(dynamicZones, ","))
}

func (d DynamicZone) EnumValues() []string {
	return dynamicZones
}

type Zone jsonschema.BasicEnum

func (z Zone) Validate(r jsonschema.RootElement) error {
	values := z.EnumValues()
	for _, v := range values {
		if v == string(z) {
			return nil
		}
	}
	return errors.New("Zone must be on of " + strings.Join(zones, ","))
}

func (z Zone) EnumValues() []string {
	return zones
}

type ActionZone jsonschema.BasicEnum

func (a ActionZone) Validate(r jsonschema.RootElement) error {
	values := a.EnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionZone must be on of " + strings.Join(actionZones, ","))
}

func (a ActionZone) EnumValues() []string {
	return actionZones
}

type EntityZone jsonschema.BasicEnum

func (e EntityZone) Validate(r jsonschema.RootElement) error {
	values := e.EnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityZone must be on of " + strings.Join(entityZones, ","))
}

func (e EntityZone) EnumValues() []string {
	return entityZones
}

type PlaceZone jsonschema.BasicEnum

func (p PlaceZone) Validate(r jsonschema.RootElement) error {
	values := p.EnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceZone must be on of " + strings.Join(placeZones, ","))
}

func (p PlaceZone) EnumValues() []string {
	return placeZones
}
