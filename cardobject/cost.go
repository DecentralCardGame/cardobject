package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Cost struct {
	ManaCost      *ManaCost      `json:",omitempty"`
	SacrificeCost *SacrificeCost `json:",omitempty"`
	DiscardCost   *DiscardCost   `json:",omitempty"`
	VoidCost      *VoidCost      `json:",omitempty"`
}

func (c Cost) ValidateType(r jsonschema.RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (c Cost) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(c)
}

type AdditionalCost struct {
	SacrificeCost *SacrificeCost `json:",omitempty"`
	DiscardCost   *DiscardCost   `json:",omitempty"`
	VoidCost      *VoidCost      `json:",omitempty"`
}

func (a AdditionalCost) ValidateType(r jsonschema.RootElement) error {
	implementer, err := a.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (a AdditionalCost) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(a)
}

type ManaCost struct {
	CostAmount BasicAmount
}

func (m ManaCost) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(m, r)
}

func (m ManaCost) InteractionText() string {
	return "§CostAmount mana"
}

type SacrificeCost struct {
	Amount BasicAmount
}

func (s SacrificeCost) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(s, r)
}

func (s SacrificeCost) InteractionText() string {
	return "Sacrifice §Amount card"
}

type DiscardCost struct {
	Amount BasicAmount
}

func (d DiscardCost) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(d, r)
}

func (d DiscardCost) InteractionText() string {
	return "Discard §Amount card"
}

type VoidCost struct {
	Amount BasicAmount
}

func (v VoidCost) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(v, r)
}

func (v VoidCost) InteractionText() string {
	return "Void §Amount card"
}
