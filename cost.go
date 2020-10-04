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
	Amount    int
	Condition *cardCondition
}

type discardCost struct {
	Amount    int
	Condition *cardCondition
}

func (c *cost) validate() error {
	possibleImplementer := []validateable{c.RessourceCost, c.SacrificeCost, c.DiscardCost}

	implementer := xorInterface(possibleImplementer)
	if implementer == nil {
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
	errors = append(errors, c.Condition.validate())
	return combineErrors(errors)
}

func (c *discardCost) validate() error {
	errors := []error{}
	errors = append(errors, validateSimpleInt(c.Amount))
	errors = append(errors, c.Condition.validate())
	return combineErrors(errors)
}
