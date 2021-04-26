package cardobject

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
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
	Class       Class
	Effects     Effects
	FlavourText FlavourText
	Tags        Tags
	Keywords    Keywords
	RulesTexts  RulesTexts
}

func (a Action) Validate() error {
	return a.ValidateStruct()
}

func (a Action) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a Action) InteractionText() string {
	return "§CardName §CastingCost §Class §Effects §FlavourText §Tags §Keywords §RulesTexts"
}

type Entity struct {
	CardName    CardName
	CastingCost CastingCost
	Class       Class
	Abilities   Abilities
	Attack      Attack
	Health      Health
	FlavourText FlavourText
	Tags        Tags
	Keywords    Keywords
	RulesTexts  RulesTexts
}

func (e Entity) Validate() error {
	return e.ValidateStruct()
}

func (e Entity) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (a Entity) InteractionText() string {
	return "§CardName §CastingCost §Class §Abilities §Attack §Health §FlavourText §Tags §Keywords §RulesTexts"
}

type Place struct {
	CardName    CardName
	CastingCost CastingCost
	Class       Class
	Abilities   Abilities
	Health      Health
	FlavourText FlavourText
	Tags        Tags
	Keywords    Keywords
	RulesTexts  RulesTexts
}

func (p Place) Validate() error {
	return p.ValidateStruct()
}

func (p Place) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (a Place) InteractionText() string {
	return "§CardName §CastingCost §Class §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts"
}

type Headquarter struct {
	CardName    CardName
	Class       Class
	Delay       Delay
	Abilities   Abilities
	Health      Health
	FlavourText FlavourText
	Tags        Tags
	Keywords    Keywords
	RulesTexts  RulesTexts
}

func (h Headquarter) Validate() error {
	return h.ValidateStruct()
}

func (h Headquarter) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (a Headquarter) InteractionText() string {
	return "§CardName §Class §Delay §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts"
}
