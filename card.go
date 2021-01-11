package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type card struct {
	*jsonschema.BasicInterface
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

type action struct {
	*jsonschema.BasicStruct
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Effects     effects
	FlavourText flavourText
	Tags        tags
}

func (a action) GetInteractionText() string {
	return "No text needed"
}

type entity struct {
	*jsonschema.BasicStruct
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Abilities   abilities
	Attack      attack
	Health      health
	FlavourText flavourText
	Tags        tags
}

func (a entity) GetInteractionText() string {
	return "No text needed"
}

type place struct {
	*jsonschema.BasicStruct
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Abilities   abilities
	Health      health
	FlavourText flavourText
	Tags        tags
}

func (a place) GetInteractionText() string {
	return "No text needed"
}

type headquarter struct {
	*jsonschema.BasicStruct
	CardName         cardName
	CostType         ressourceCostType
	Abilities        abilities
	Health           health
	Growth           growth
	StartingHandSize startingHandsize
	Wisdom           wisdom
	FlavourText      flavourText
	Tags             tags
}

func (a headquarter) GetInteractionText() string {
	return "No text needed"
}
