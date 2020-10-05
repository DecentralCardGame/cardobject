package cardobject

import "errors"

type cost struct {
	RessourceCost *ressourceCost `json:",omitempty"`
	SacrificeCost *sacrificeCost `json:",omitempty"`
	DiscardCost   *discardCost   `json:",omitempty"`
}

type ressourceCost struct {
	RessourceCostType *ressourceCostType `json:",omitempty"`
	CostAmount        int
}

type sacrificeCost struct {
	Amount     int
	Conditions *cardConditions
}

type discardCost struct {
	Amount     int
	Conditions *cardConditions
}

func (c *cost) validate() error {
	possibleImplementer := []validateable{c.RessourceCost, c.SacrificeCost, c.DiscardCost}

	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New("Ability implemented by not exactly one option")
	}
	return implementer.validate()
}

func (c *ressourceCost) validate() error {
	errors := []error{}
	errors = append(errors, validateCastingCost(c.CostAmount))
	errors = append(errors, c.RessourceCostType.validate())
	return combineErrors(errors)
}

func (c *sacrificeCost) validate() error {
	errors := []error{}
	errors = append(errors, validateSimpleInt(c.Amount))
	if c.Conditions != nil {
		errors = append(errors, c.Conditions.validate())
	}
	return combineErrors(errors)
}

func (c *discardCost) validate() error {
	errors := []error{}
	errors = append(errors, validateSimpleInt(c.Amount))
	if c.Conditions != nil {
		errors = append(errors, c.Conditions.validate())
	}
	return combineErrors(errors)
}
