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

func (c Cost) Validate() error {
	return c.ValidateStruct()
}

func (c Cost) ValidateStruct() error {
	return jsonschema.ValidateStruct(c)
}

func (c Cost) InteractionText() string {
	return "§ManaCost §SacrificeCost §DiscardCost §VoidCost"
}

type AdditionalCost struct {
	SacrificeCost *SacrificeCost `json:",omitempty"`
	DiscardCost   *DiscardCost   `json:",omitempty"`
	VoidCost      *VoidCost      `json:",omitempty"`
}

func (a AdditionalCost) Validate() error {
	return a.ValidateStruct()
}

func (a AdditionalCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a AdditionalCost) InteractionText() string {
	return "§SacrificeCost §DiscardCost §VoidCost"
}

type ManaCost struct {
	CostAmount BasicAmount
}

func (m ManaCost) Validate() error {
	return m.ValidateStruct()
}

func (m ManaCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(m)
}

func (m ManaCost) InteractionText() string {
	return "§CostAmount mana"
}

type SacrificeCost struct {
	Amount BasicAmount
}

func (s SacrificeCost) Validate() error {
	return s.ValidateStruct()
}

func (s SacrificeCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(s)
}

func (s SacrificeCost) InteractionText() string {
	return "Sacrifice §Amount card"
}

type DiscardCost struct {
	Amount BasicAmount
}

func (d DiscardCost) Validate() error {
	return d.ValidateStruct()
}

func (d DiscardCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(d)
}

func (d DiscardCost) InteractionText() string {
	return "Discard §Amount card"
}

type VoidCost struct {
	Amount BasicAmount
}

func (v VoidCost) Validate() error {
	return v.ValidateStruct()
}

func (v VoidCost) ValidateStruct() error {
	return jsonschema.ValidateStruct(v)
}

func (v VoidCost) InteractionText() string {
	return "Void §Amount card"
}
