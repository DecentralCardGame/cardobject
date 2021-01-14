package cardobject

import (
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
	return ""
}

type sacrificeCost struct {
	*jsonschema.BasicStruct
	Amount     basicAmount
	Conditions *cardConditions
}

func (s sacrificeCost) GetInteractionText() string {
	return ""
}

type discardCost struct {
	*jsonschema.BasicStruct
	Amount     basicAmount
	Conditions *cardConditions
}

func (d discardCost) GetInteractionText() string {
	return ""
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
	return "No InteractionText defined"
}

type energy struct {
	*jsonschema.BasicBool
}

type food struct {
	*jsonschema.BasicBool
}
type lumber struct {
	*jsonschema.BasicBool
}
type mana struct {
	*jsonschema.BasicBool
}
type iron struct {
	*jsonschema.BasicBool
}
