package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type card struct {
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

func (c card) Validate() error {
	return c.ValidateInterface()
}

func (c card) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type action struct {
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Effects     effects
	FlavourText flavourText
	Tags        tags
}

func (a action) Validate() error {
	return a.ValidateStruct()
}

func (a action) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a action) GetInteractionText() string {
	return "§CardName §CastingCost §CostType §Effects §FlavourText §Tags"
}

type entity struct {
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Abilities   abilities
	Attack      attack
	Health      health
	FlavourText flavourText
	Tags        tags
}

func (e entity) Validate() error {
	return e.ValidateStruct()
}

func (e entity) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (a entity) GetInteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Attack §Health §FlavourText §Tags"
}

type place struct {
	CardName    cardName
	CastingCost castingCost
	CostType    ressourceCostType
	Abilities   abilities
	Health      health
	FlavourText flavourText
	Tags        tags
}

func (p place) Validate() error {
	return p.ValidateStruct()
}

func (p place) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (a place) GetInteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Health §FlavourText §Tags"
}

type headquarter struct {
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

func (h headquarter) Validate() error {
	return h.ValidateStruct()
}

func (h headquarter) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (a headquarter) GetInteractionText() string {
	return "§CardName §CostType §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags"
}
