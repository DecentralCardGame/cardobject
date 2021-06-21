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

func (c Card) ValidateClasses(s []jsonschema.Class) error {
	return nil
}

func (c Card) Validate() error {
	return c.ValidateType(c)
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

type Action struct {
	CardName       CardName
	CastingCost    CastingCost
	AdditionalCost *AdditionalCost `json:",omitempty"`
	Class          Class
	Effects        Effects
	FlavourText    FlavourText
	Tags           Tags
	Keywords       Keywords
	RulesTexts     RulesTexts
}

func (a Action) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a Action) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Effects §FlavourText §Tags §Keywords §RulesTexts"
}

type Entity struct {
	CardName       CardName
	CastingCost    CastingCost
	AdditionalCost *AdditionalCost `json:",omitempty"`
	Class          Class
	Abilities      Abilities
	Attack         Attack
	Health         Health
	FlavourText    FlavourText
	Tags           Tags
	Keywords       Keywords
	RulesTexts     RulesTexts
}

func (e Entity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (a Entity) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Attack §Health §FlavourText §Tags §Keywords §RulesTexts"
}

type Place struct {
	CardName       CardName
	CastingCost    CastingCost
	AdditionalCost *AdditionalCost `json:",omitempty"`
	Class          Class
	Abilities      Abilities
	Health         Health
	FlavourText    FlavourText
	Tags           Tags
	Keywords       Keywords
	RulesTexts     RulesTexts
}

func (p Place) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (a Place) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts"
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

func (h Headquarter) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (a Headquarter) InteractionText() string {
	return "§CardName §Class §Delay §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts"
}
