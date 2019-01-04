package cardobject

import "testing"
import "reflect"

func TestActivatedAbility(t *testing.T) {
	testedInterface := "ActivatedAbility"
	x := createActivatedAbility()
	if(!reflect.DeepEqual(x.GetEffect(), createEffect())) {
		t.Errorf(testedInterface + " doesn't return Effect")
	}
	if(!reflect.DeepEqual(x.GetCost(), createCost())) {
		t.Errorf(testedInterface + " doesn't return Cost")
	}
	if(x.IsMultipleUse() != false) {
		t.Errorf(testedInterface + " doesn't return MultipleUse")
	}
	success(testedInterface)
}

func TestTriggeredAbility(t *testing.T) {
	testedInterface := "TriggeredAbility"
	x := createTriggeredAbility()
	if(!reflect.DeepEqual(x.GetEffect(), createEffect())) {
		t.Errorf(testedInterface + " doesn't return Effect")
	}
	if(!reflect.DeepEqual(x.GetEventListener(), createTimeEventListener())) {
		t.Errorf(testedInterface + " doesn't return EventListener")
	}
	success(testedInterface)
}

func createActivatedAbility() ActivatedAbility {
	return NewActivatedAbility(createSpeedModifier(), createEffect(), createCost(), false)
}

func createTriggeredAbility() TriggeredAbility {
	return NewTriggeredAbility(createSpeedModifier(), createEffect(), createTimeEventListener())
}
