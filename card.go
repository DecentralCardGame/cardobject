package cardobject

import (
	"errors"
)

type card struct {
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

type action struct {
	CardName    string
	CastingCost int
	CostType    *ressourceCostType
	Effects     []effect
	FlavourText string
	Tags        []string
}

type entity struct {
	CardName    string
	CastingCost int
	CostType    *ressourceCostType
	Abilities   []ability
	Attack      int
	Health      int
	FlavourText string
	Tags        []string
}

type place struct {
	CardName    string
	CastingCost int
	CostType    *ressourceCostType
	Abilities   []ability
	Health      int
	FlavourText string
	Tags        []string
}

type headquarter struct {
	CardName         string
	CostType         *ressourceCostType
	Abilities        []ability
	Health           int
	Growth           int
	StartingHandSize int
	Wisdom           int
	FlavourText      string
	Tags             []string
}

func (c *card) validate() error {
	possibleImplementer := []validateable{c.Action, c.Entity, c.Place, c.Headquarter}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Card implemented by not exactly one option")
	}
	return implementer.validate()
}

func (a *action) validate() error {
	errors := []error{}
	errors = append(errors, validateCardName(a.CardName))
	errors = append(errors, validateCastingCost(a.CastingCost))
	errors = append(errors, a.CostType.validate())
	errors = append(errors, validateEffects(a.Effects))
	errors = append(errors, validateFlavourText(a.FlavourText))
	errors = append(errors, validateTags(a.Tags))
	return combineErrors(errors)
}

func (e *entity) validate() error {
	errors := []error{}
	errors = append(errors, validateCardName(e.CardName))
	errors = append(errors, validateCastingCost(e.CastingCost))
	errors = append(errors, e.CostType.validate())
	errors = append(errors, validateAbilities(e.Abilities))
	errors = append(errors, validateAttack(e.Attack))
	errors = append(errors, validateHealth(e.Health))
	errors = append(errors, validateFlavourText(e.FlavourText))
	errors = append(errors, validateTags(e.Tags))
	return combineErrors(errors)
}

func (p *place) validate() error {
	errors := []error{}
	errors = append(errors, validateCardName(p.CardName))
	errors = append(errors, validateCastingCost(p.CastingCost))
	errors = append(errors, p.CostType.validate())
	errors = append(errors, validateAbilities(p.Abilities))
	errors = append(errors, validateHealth(p.Health))
	errors = append(errors, validateFlavourText(p.FlavourText))
	errors = append(errors, validateTags(p.Tags))
	return combineErrors(errors)
}

func (h *headquarter) validate() error {
	errors := []error{}
	errors = append(errors, validateCardName(h.CardName))
	errors = append(errors, h.CostType.validate())
	errors = append(errors, validateAbilities(h.Abilities))
	errors = append(errors, validateGrowth(h.Growth))
	errors = append(errors, validateStartingHandSize(h.StartingHandSize))
	errors = append(errors, validateWisdom(h.Wisdom))
	errors = append(errors, validateHealth(h.Health))
	errors = append(errors, validateFlavourText(h.FlavourText))
	errors = append(errors, validateTags(h.Tags))
	return combineErrors(errors)
}
