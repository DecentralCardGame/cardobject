package keywords

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type Card struct {
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

func (c Card) Validate() error {
	return c.ValidateInterface()
}

func (c Card) ValidateInterface() error {
	return jsonschema.ValidateInterface(c)
}

type action struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Effects        effects
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	RulesTexts     cardobject.RulesTexts
}

func (a action) Validate() error {
	return a.ValidateStruct()
}

func (a action) ValidateStruct() error {
	return jsonschema.ValidateStruct(a)
}

func (a action) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Effects §FlavourText §Tags §Keywords §RulesTexts"
}

type entity struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Abilities      abilities
	Attack         cardobject.Attack
	Health         cardobject.Health
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	RulesTexts     cardobject.RulesTexts
}

func (e entity) Validate() error {
	return e.ValidateStruct()
}

func (e entity) ValidateStruct() error {
	return jsonschema.ValidateStruct(e)
}

func (e entity) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Attack §Health §FlavourText §Tags §Keywords §RulesTexts"
}

type place struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Abilities      abilities
	Health         cardobject.Health
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	RulesTexts     cardobject.RulesTexts
}

func (p place) Validate() error {
	return p.ValidateStruct()
}

func (p place) ValidateStruct() error {
	return jsonschema.ValidateStruct(p)
}

func (p place) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts"
}

type headquarter struct {
	CardName    cardobject.CardName
	Class       cardobject.Class
	Delay       cardobject.Delay
	Abilities   abilities
	Health      cardobject.Health
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
	Keywords    cardobject.Keywords
	RulesTexts  cardobject.RulesTexts
}

func (h headquarter) Validate() error {
	return h.ValidateStruct()
}

func (h headquarter) ValidateStruct() error {
	return jsonschema.ValidateStruct(h)
}

func (h headquarter) InteractionText() string {
	return "§CardName §Class §Delay §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags §Keywords §RulesTexts"
}
