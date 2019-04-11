package cardobject

import "testing"
import "reflect"

func TestCost(t *testing.T) {
	testedInterface := "Cost"
	x := createCost()
	if(!reflect.DeepEqual(x.GetRessources(), []Ressource{MANA})) {
		t.Errorf(testedInterface + " doesn't return Ressources")
	}
	success(testedInterface)
}

func createCost() Cost {
	return NewCost([]Ressource{MANA})
}