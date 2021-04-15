package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Class struct {
	Nature     Nature
	Mysticism  Mysticism
	Technology Technology
	Culture    Culture
}

func (c Class) Validate() error {
	if !(bool(c.Nature) || bool(c.Mysticism) || bool(c.Technology) || bool(c.Culture)) {
		return errors.New("At least one Class Must be selected for a card.")
	}
	return c.ValidateStruct()
}

func (c Class) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c Class) InteractionText() string {
	return "§Nature §Mysticism §Technology §Culture"
}

type Nature jsonschema.BasicBool

func (n Nature) Validate() error {
	return n.ValidateBool()
}

func (n Nature) ValidateBool() error {
	b := bool(n)
	if b || !b {
		return nil
	}
	return errors.New("Nature must be true or false")
}

type Mysticism jsonschema.BasicBool

func (m Mysticism) Validate() error {
	return m.ValidateBool()
}

func (m Mysticism) ValidateBool() error {
	b := bool(m)
	if b || !b {
		return nil
	}
	return errors.New("Mysticism must be true or false")
}

type Technology jsonschema.BasicBool

func (t Technology) Validate() error {
	return t.ValidateBool()
}

func (t Technology) ValidateBool() error {
	b := bool(t)
	if b || !b {
		return nil
	}
	return errors.New("Technology must be true or false")
}

type Culture jsonschema.BasicBool

func (c Culture) Validate() error {
	return c.ValidateBool()
}

func (c Culture) ValidateBool() error {
	b := bool(c)
	if b || !b {
		return nil
	}
	return errors.New("Culture must be true or false")
}
