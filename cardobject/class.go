package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//CULTURE Class
const CULTURE jsonschema.Class = "CULTURE"

//NATURE Class
const NATURE jsonschema.Class = "NATURE"

//MYSTICISM Class
const MYSTICISM jsonschema.Class = "MYSTICISM"

//TECHNOLOGY Class
const TECHNOLOGY jsonschema.Class = "TECHNOLOGY"

type Class struct {
	Nature     Nature
	Mysticism  Mysticism
	Technology Technology
	Culture    Culture
}

func (c Class) ValidateType(r jsonschema.RootElement) error {
	if !(bool(c.Nature) || bool(c.Mysticism) || bool(c.Technology) || bool(c.Culture)) {
		return errors.New("At least one Class Must be selected for a card.")
	}
	return jsonschema.ValidateStruct(c, r)
}

func (c Class) InteractionText() string {
	return "§Nature §Mysticism §Technology §Culture"
}

func (c Class) contains(s []jsonschema.Class) bool {
	for _, v := range s {
		if bool(c.Culture) && (v == CULTURE) || bool(c.Mysticism) && (v == MYSTICISM) ||
			bool(c.Nature) && (v == NATURE) || bool(c.Technology) && (v == TECHNOLOGY) {
			return true
		}
	}
	return false
}

type Nature jsonschema.BasicBool

func (n Nature) ValidateType(r jsonschema.RootElement) error {
	b := bool(n)
	if b || !b {
		return nil
	}
	return errors.New("Nature must be true or false")
}

func (n Nature) Default() bool {
	return false
}

type Mysticism jsonschema.BasicBool

func (m Mysticism) ValidateType(r jsonschema.RootElement) error {
	b := bool(m)
	if b || !b {
		return nil
	}
	return errors.New("Mysticism must be true or false")
}

func (m Mysticism) Default() bool {
	return false
}

type Technology jsonschema.BasicBool

func (t Technology) ValidateType(r jsonschema.RootElement) error {
	b := bool(t)
	if b || !b {
		return nil
	}
	return errors.New("Technology must be true or false")
}

func (t Technology) Default() bool {
	return false
}

type Culture jsonschema.BasicBool

func (c Culture) ValidateType(r jsonschema.RootElement) error {
	b := bool(c)
	if b || !b {
		return nil
	}
	return errors.New("Culture must be true or false")
}

func (c Culture) Default() bool {
	return false
}
