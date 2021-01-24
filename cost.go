package cardobject

import (
	"errors"

	"github.com/DecentralCardGame/jsonschema"
)

type cost struct {
	*jsonschema.BasicInterface
	RessourceCost *ressourceCost `json:",omitempty"`
	SacrificeCost *sacrificeCost `json:",omitempty"`
	DiscardCost   *discardCost   `json:",omitempty"`
}

type ressourceCost struct {
	*jsonschema.BasicStruct
	CostAmount basicAmount
}

func (r ressourceCost) GetInteractionText() string {
	return "§CostAmount ressources"
}

type sacrificeCost struct {
	*jsonschema.BasicStruct
	Amount     basicAmount
	Conditions *cardConditions
}

func (s sacrificeCost) GetInteractionText() string {
	return "Sacrifice §Amount card §Conditions"
}

type discardCost struct {
	*jsonschema.BasicStruct
	Amount     basicAmount
	Conditions *cardConditions
}

func (d discardCost) GetInteractionText() string {
	return "Discard §Amount card §Conditions"
}

type ressourceCostType struct {
	*jsonschema.BasicStruct
	Energy energy
	Food   food
	Lumber lumber
	Mana   mana
	Iron   iron
}

func (r ressourceCostType) GetInteractionText() string {
	return "§Energy §Food §Lumber §Mana §Iron"
}

type energy jsonschema.BasicBool

func (e energy) Validate() error {
	return e.ValidateBool()
}

func (e energy) ValidateBool() error {
	b := bool(e)
	if b || !b {
		return nil
	}
	return errors.New("")
}

type food jsonschema.BasicBool

func (f food) Validate() error {
	return f.ValidateBool()
}

func (f food) ValidateBool() error {
	b := bool(f)
	if b || !b {
		return nil
	}
	return errors.New("")
}

type lumber jsonschema.BasicBool

func (l lumber) Validate() error {
	return l.ValidateBool()
}

func (l lumber) ValidateBool() error {
	b := bool(l)
	if b || !b {
		return nil
	}
	return errors.New("")
}

type mana jsonschema.BasicBool

func (m mana) Validate() error {
	return m.ValidateBool()
}

func (m mana) ValidateBool() error {
	b := bool(m)
	if b || !b {
		return nil
	}
	return errors.New("")
}

type iron jsonschema.BasicBool

func (i iron) Validate() error {
	return i.ValidateBool()
}

func (i iron) ValidateBool() error {
	b := bool(i)
	if b || !b {
		return nil
	}
	return errors.New("")
}
