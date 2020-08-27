package cardobject

import "errors"

type ressourceCostType struct {
	Energy bool
	Food   bool
	Lumber bool
	Mana   bool
	Iron   bool
}

func (costType *ressourceCostType) validate() error {
	if costType == nil {
		return errors.New("The cost type is nil")
	}
	return nil
}
