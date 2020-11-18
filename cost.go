package cardobject

import "errors"

type cost struct {
	RessourceCost *ressourceCost `json:",omitempty"`
	SacrificeCost *sacrificeCost `json:",omitempty"`
	DiscardCost   *discardCost   `json:",omitempty"`
}

type ressourceCost struct {
	CostAmount int
}

type sacrificeCost struct {
	Amount     int
	Conditions *cardConditions
}

type discardCost struct {
	Amount     int
	Conditions *cardConditions
}

type ressourceCostType struct {
	Energy bool
	Food   bool
	Lumber bool
	Mana   bool
	Iron   bool
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
	errorRange := []error{}
	errorRange = append(errorRange, validateCastingCost(c.CostAmount))
	return combineErrors(errorRange)
}

func (c *sacrificeCost) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateSimpleInt(c.Amount))
	if c.Conditions != nil {
		errorRange = append(errorRange, c.Conditions.validate())
	}
	return combineErrors(errorRange)
}

func (c *discardCost) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateSimpleInt(c.Amount))
	if c.Conditions != nil {
		errorRange = append(errorRange, c.Conditions.validate())
	}
	return combineErrors(errorRange)
}
