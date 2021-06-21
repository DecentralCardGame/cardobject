package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

//CULTURE Class
const CULTURE = "CULTURE"

//NATURE Class
const NATURE = "NATURE"

//MYSTICISM Class
const MYSTICISM = "MYSTICISM"

//TECHNOLOGY Class
const TECHNOLOGY = "TECHNOLOGY"

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

func (c Class) Contains(s string) bool {
	b := bool(c.Culture) && (s == CULTURE) || bool(c.Mysticism) && (s == MYSTICISM) ||
		bool(c.Nature) && (s == NATURE) || bool(c.Technology) && (s == TECHNOLOGY)
	return b
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
