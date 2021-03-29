package cardobject

import (
	"errors"
	"strings"

	"github.com/DecentralCardGame/jsonschema"
)

var dynamicZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND"}
var zones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

var actionZones []string = []string{"DECK", "DUSTPILE", "HAND", "VOID"}
var entityZones []string = []string{"ATTACKLANE", "BLOCKLANE", "DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}
var placeZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

type DynamicZone jsonschema.BasicEnum

func (d DynamicZone) Validate() error {
	return d.ValidateEnum()
}

func (d DynamicZone) ValidateEnum() error {
	values := d.GetEnumValues()
	for _, v := range values {
		if v == string(d) {
			return nil
		}
	}
	return errors.New("DynamicZone must be on of " + strings.Join(dynamicZones, ","))
}

func (d DynamicZone) GetEnumValues() []string {
	return dynamicZones
}

type Zone jsonschema.BasicEnum

func (z Zone) Validate() error {
	return z.ValidateEnum()
}

func (z Zone) ValidateEnum() error {
	values := z.GetEnumValues()
	for _, v := range values {
		if v == string(z) {
			return nil
		}
	}
	return errors.New("Zone must be on of " + strings.Join(zones, ","))
}

func (z Zone) GetEnumValues() []string {
	return zones
}

type ActionZone jsonschema.BasicEnum

func (a ActionZone) Validate() error {
	return a.ValidateEnum()
}

func (a ActionZone) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("ActionZone must be on of " + strings.Join(actionZones, ","))
}

func (a ActionZone) GetEnumValues() []string {
	return actionZones
}

type EntityZone jsonschema.BasicEnum

func (e EntityZone) Validate() error {
	return e.ValidateEnum()
}

func (e EntityZone) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("EntityZone must be on of " + strings.Join(entityZones, ","))
}

func (e EntityZone) GetEnumValues() []string {
	return entityZones
}

type PlaceZone jsonschema.BasicEnum

func (p PlaceZone) Validate() error {
	return p.ValidateEnum()
}

func (p PlaceZone) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("PlaceZone must be on of " + strings.Join(placeZones, ","))
}

func (p PlaceZone) GetEnumValues() []string {
	return placeZones
}
