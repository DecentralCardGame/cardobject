package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

var dynamicZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND"}
var zones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

var actionZones []string = []string{"DECK", "DUSTPILE", "HAND", "VOID"}
var entityZones []string = []string{"ATTACKLANE", "BLOCKLANE", "DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}
var placeZones []string = []string{"DECK", "DUSTPILE", "FIELD", "HAND", "VOID"}

type dynamicZone jsonschema.BasicEnum

func (d dynamicZone) Validate() error {
	return d.ValidateEnum()
}

func (d dynamicZone) ValidateEnum() error {
	values := d.GetEnumValues()
	for _, v := range values {
		if v == string(d) {
			return nil
		}
	}
	return errors.New("")
}

func (d dynamicZone) GetEnumValues() []string {
	return dynamicZones
}

type zone jsonschema.BasicEnum

func (z zone) Validate() error {
	return z.ValidateEnum()
}

func (z zone) ValidateEnum() error {
	values := z.GetEnumValues()
	for _, v := range values {
		if v == string(z) {
			return nil
		}
	}
	return errors.New("")
}

func (z zone) GetEnumValues() []string {
	return zones
}

type actionZone jsonschema.BasicEnum

func (a actionZone) Validate() error {
	return a.ValidateEnum()
}

func (a actionZone) ValidateEnum() error {
	values := a.GetEnumValues()
	for _, v := range values {
		if v == string(a) {
			return nil
		}
	}
	return errors.New("")
}

func (a actionZone) GetEnumValues() []string {
	return actionZones
}

type entityZone jsonschema.BasicEnum

func (e entityZone) Validate() error {
	return e.ValidateEnum()
}

func (e entityZone) ValidateEnum() error {
	values := e.GetEnumValues()
	for _, v := range values {
		if v == string(e) {
			return nil
		}
	}
	return errors.New("")
}

func (e entityZone) GetEnumValues() []string {
	return entityZones
}

type placeZone jsonschema.BasicEnum

func (p placeZone) Validate() error {
	return p.ValidateEnum()
}

func (p placeZone) ValidateEnum() error {
	values := p.GetEnumValues()
	for _, v := range values {
		if v == string(p) {
			return nil
		}
	}
	return errors.New("")
}

func (p placeZone) GetEnumValues() []string {
	return placeZones
}
