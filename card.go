package cardobject

import "errors"

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
	Effects     effects
	FlavourText string
	Tags        []string
}

type entity struct {
	CardName    string
	CastingCost int
	CostType    *ressourceCostType
	Abilities   abilities
	Attack      int
	Health      int
	FlavourText string
	Tags        []string
}

type place struct {
	CardName    string
	CastingCost int
	CostType    *ressourceCostType
	Abilities   abilities
	Health      int
	FlavourText string
	Tags        []string
}

type headquarter struct {
	CardName         string
	CostType         *ressourceCostType
	Abilities        abilities
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
	errorRange := []error{}
	errorRange = append(errorRange, validateCardName(a.CardName))
	errorRange = append(errorRange, validateCastingCost(a.CastingCost))
	if a.CostType == nil {
		errorRange = append(errorRange, errors.New("Action must have a CostType"))
	}
	errorRange = append(errorRange, a.Effects.validate())
	errorRange = append(errorRange, validateFlavourText(a.FlavourText))
	errorRange = append(errorRange, validateTags(a.Tags))
	return combineErrors(errorRange)
}

func (e *entity) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardName(e.CardName))
	errorRange = append(errorRange, validateCastingCost(e.CastingCost))
	if e.CostType == nil {
		errorRange = append(errorRange, errors.New("Entity must have a CostType"))
	}
	errorRange = append(errorRange, e.Abilities.validate())
	errorRange = append(errorRange, validateAttack(e.Attack))
	errorRange = append(errorRange, validateHealth(e.Health))
	errorRange = append(errorRange, validateFlavourText(e.FlavourText))
	errorRange = append(errorRange, validateTags(e.Tags))
	return combineErrors(errorRange)
}

func (p *place) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardName(p.CardName))
	errorRange = append(errorRange, validateCastingCost(p.CastingCost))
	if p.CostType == nil {
		errorRange = append(errorRange, errors.New("Place must have a CostType"))
	}
	errorRange = append(errorRange, p.Abilities.validate())
	errorRange = append(errorRange, validateHealth(p.Health))
	errorRange = append(errorRange, validateFlavourText(p.FlavourText))
	errorRange = append(errorRange, validateTags(p.Tags))
	return combineErrors(errorRange)
}

func (h *headquarter) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardName(h.CardName))
	if h.CostType == nil {
		errorRange = append(errorRange, errors.New("HQ must have a CostType"))
	}
	errorRange = append(errorRange, h.Abilities.validate())
	errorRange = append(errorRange, validateGrowth(h.Growth))
	errorRange = append(errorRange, validateStartingHandSize(h.StartingHandSize))
	errorRange = append(errorRange, validateWisdom(h.Wisdom))
	errorRange = append(errorRange, validateHealth(h.Health))
	errorRange = append(errorRange, validateFlavourText(h.FlavourText))
	errorRange = append(errorRange, validateTags(h.Tags))
	return combineErrors(errorRange)
}
