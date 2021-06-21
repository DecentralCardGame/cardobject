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
	return c.ValidateType(c)
}

func (c Card) ValidateClasses(s []string) error {
	return nil
}

func (c Card) ValidateType(r jsonschema.RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (c Card) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(c)
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

func (a action) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
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

func (e entity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
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

func (p place) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
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

func (h headquarter) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h headquarter) InteractionText() string {
	return "§CardName §Class §Delay §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags §Keywords §RulesTexts"
}
