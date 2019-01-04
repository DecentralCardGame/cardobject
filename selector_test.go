package cardobject

import "testing"
import "reflect"


func TestCardSelectorCond(t *testing.T) {
	testedInterface := "CardSelectorCond"
	x := createBasicSelector()
	if(x.GetPlayerSelectorMode() != ALL) {
		t.Errorf(testedInterface + " doesn't return PlayerSelectorMode")
	}
	if(!reflect.DeepEqual(x.GetPlayerCondition(), createPlayerIntCondition())) {
		t.Errorf(testedInterface + " doesn't return PlayerCondition")
	}
	if(x.GetCardSelectorMode() != ALL) {
		t.Errorf(testedInterface + " doesn't return CardSelectorMode")
	}
	if(!reflect.DeepEqual(x.GetCardCondition(), createCardIntCondition())) {
		t.Errorf(testedInterface + " doesn't return CardCondition")
	}
	if(x.GetZone() != EXILE) {
		t.Errorf(testedInterface + " doesn't return Zone")
	}
	success(testedInterface)
}

func TestCardSelectorSafeCond(t *testing.T) {
	testedInterface := "CardSelectorSafeCond"
	x := createSafeSelector()
	if(x.GetPlayerSelectorMode() != ALL) {
		t.Errorf(testedInterface + " doesn't return PlayerSelectorMode")
	}
	if(!reflect.DeepEqual(x.GetPlayerCondition(), createPlayerIntCondition())) {
		t.Errorf(testedInterface + " doesn't return PlayerCondition")
	}
	if(x.GetCardSelectorMode() != ALL) {
		t.Errorf(testedInterface + " doesn't return CardSelectorMode")
	}
	if(!reflect.DeepEqual(x.GetCardCondition(), createCardIntCondition())) {
		t.Errorf(testedInterface + " doesn't return CardCondition")
	}
	if(x.GetZone() != HAND) {
		t.Errorf(testedInterface + " doesn't return Zone")
	}
	if(x.GetDynamicZone() != HAND) {
		t.Errorf(testedInterface + " doesn't return Zone")
	}
	success(testedInterface)
}


func createBasicSelector() CardSelectorCond {
	return NewBasicSelector(ALL, createPlayerIntCondition(), ALL, createCardIntCondition(), EXILE)
}

func createSafeSelector() CardSelectorSafeCond {
	return NewSafeSelector(ALL, createPlayerIntCondition(), ALL, createCardIntCondition(), HAND)
}