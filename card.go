package cardobject

import (
	"github.com/DecentralCardGame/jsonschema"
)

type Card struct {
	Action      *Action      `json:",omitempty"`
	Entity      *Entity      `json:",omitempty"`
	Place       *Place       `json:",omitempty"`
	Headquarter *Headquarter `json:",omitempty"`
}

func (c Card) Validate() error {
	return c.ValidateInterface()
}

func (c Card) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type Action struct {
	CardName    CardName
	CastingCost CastingCost
	CostType    RessourceCostType
	Effects     Effects
	FlavourText FlavourText
	Tags        Tags
}

func (a Action) Validate() error {
	return a.ValidateStruct()
}

func (a Action) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a Action) InteractionText() string {
	return "§CardName §CastingCost §CostType §Effects §FlavourText §Tags"
}

type Entity struct {
	CardName    CardName
	CastingCost CastingCost
	CostType    RessourceCostType
	Abilities   Abilities
	Attack      Attack
	Health      Health
	FlavourText FlavourText
	Tags        Tags
}

func (e Entity) Validate() error {
	return e.ValidateStruct()
}

func (e Entity) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (a Entity) InteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Attack §Health §FlavourText §Tags"
}

type Place struct {
	CardName    CardName
	CastingCost CastingCost
	CostType    RessourceCostType
	Abilities   Abilities
	Health      Health
	FlavourText FlavourText
	Tags        Tags
}

func (p Place) Validate() error {
	return p.ValidateStruct()
}

func (p Place) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (a Place) InteractionText() string {
	return "§CardName §CastingCost §CostType §Abilities §Health §FlavourText §Tags"
}

type Headquarter struct {
	CardName         CardName
	CostType         RessourceCostType
	Abilities        Abilities
	Health           Health
	Growth           Growth
	StartingHandSize StartingHandsize
	Wisdom           Wisdom
	FlavourText      FlavourText
	Tags             Tags
}

func (h Headquarter) Validate() error {
	return h.ValidateStruct()
}

func (h Headquarter) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (a Headquarter) InteractionText() string {
	return "§CardName §CostType §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags"
}
