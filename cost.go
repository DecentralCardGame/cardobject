package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

type Cost struct {
	RessourceCost *RessourceCost `json:",omitempty"`
	SacrificeCost *SacrificeCost `json:",omitempty"`
	DiscardCost   *DiscardCost   `json:",omitempty"`
}

func (c Cost) Validate() error {
	return c.ValidateInterface()
}

func (c Cost) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type RessourceCost struct {
	CostAmount BasicAmount
}

func (r RessourceCost) Validate() error {
	return r.ValidateStruct()
}

func (r RessourceCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r RessourceCost) GetInteractionText() string {
	return "§CostAmount ressources"
}

type SacrificeCost struct {
	Amount     BasicAmount
	Conditions *CardConditions
}

func (s SacrificeCost) Validate() error {
	return s.ValidateStruct()
}

func (s SacrificeCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s SacrificeCost) GetInteractionText() string {
	return "Sacrifice §Amount card §Conditions"
}

type DiscardCost struct {
	Amount     BasicAmount
	Conditions *CardConditions
}

func (d DiscardCost) Validate() error {
	return d.ValidateStruct()
}

func (d DiscardCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d DiscardCost) GetInteractionText() string {
	return "Discard §Amount card §Conditions"
}

type RessourceCostType struct {
	Energy Energy
	Food   Food
	Lumber Lumber
	Mana   Mana
	Iron   Iron
}

func (r RessourceCostType) Validate() error {
	if !(bool(r.Energy) || bool(r.Food) || bool(r.Lumber) || bool(r.Mana) || bool(r.Iron)) {
		return errors.New("At least one Costtype Must be selected for a card.")
	}
	return r.ValidateStruct()
}

func (r RessourceCostType) ValidateStruct() error {
	return jsonschema.ValidateStruct(r)
}

func (r RessourceCostType) GetInteractionText() string {
	return "§Energy §Food §Lumber §Mana §Iron"
}

type Energy jsonschema.BasicBool

func (e Energy) Validate() error {
	return e.ValidateBool()
}

func (e Energy) ValidateBool() error {
	b := bool(e)
	if b || !b {
		return nil
	}
	return errors.New("Energy must be true or false")
}

type Food jsonschema.BasicBool

func (f Food) Validate() error {
	return f.ValidateBool()
}

func (f Food) ValidateBool() error {
	b := bool(f)
	if b || !b {
		return nil
	}
	return errors.New("Food must be true or false")
}

type Lumber jsonschema.BasicBool

func (l Lumber) Validate() error {
	return l.ValidateBool()
}

func (l Lumber) ValidateBool() error {
	b := bool(l)
	if b || !b {
		return nil
	}
	return errors.New("Lumber must be true or false")
}

type Mana jsonschema.BasicBool

func (m Mana) Validate() error {
	return m.ValidateBool()
}

func (m Mana) ValidateBool() error {
	b := bool(m)
	if b || !b {
		return nil
	}
	return errors.New("Mana must be true or false")
}

type Iron jsonschema.BasicBool

func (i Iron) Validate() error {
	return i.ValidateBool()
}

func (i Iron) ValidateBool() error {
	b := bool(i)
	if b || !b {
		return nil
	}
	return errors.New("Iron must be true or false")
}
