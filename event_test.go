package cardobject

import "testing"

func TestTimeEventListener(t *testing.T) {
	testedInterface := "TimeEventListener"
	x := createTimeEventListener()
	if(!x.IsEventListener()) {
		t.Errorf(testedInterface + " isn't an eventListener")
	}
	if(x.GetTimeEvent() != TICKEND) {
		t.Errorf(testedInterface + " doesn't return TimeEvent")
	}
	success(testedInterface)
}

func TestManipulationEventListener(t *testing.T) {
	testedInterface := "ManipulationEventListener"
	x := createManipulationEventListener()
	if(!x.IsEventListener()) {
		t.Errorf(testedInterface + " isn't an eventListener")
	}
	if(x.GetPropertyId() != ATTACK) {
		t.Errorf(testedInterface + " doesn't return PropertyId")
	}
	success(testedInterface)
}

func TestZoneChangeEventListener(t *testing.T) {
	testedInterface := "ZoneChangeEventListener"
	x := createZoneChangeEventListener()
	if(!x.IsEventListener()) {
		t.Errorf(testedInterface + " isn't an eventListener")
	}
	if(x.GetSource() != HAND) {
		t.Errorf(testedInterface + " doesn't return Source")
	}
	if(x.GetDestination() != EXILE) {
		t.Errorf(testedInterface + " doesn't return Destination")
	}
	success(testedInterface)
}

func createTimeEventListener() TimeEventListener {
	return NewTimeEventListener(TICKEND)
}

func createManipulationEventListener() ManipulationEventListener {
	return NewManipulationEventListener(ATTACK)
}

func createZoneChangeEventListener() ZoneChangeEventListener {
	return NewZoneChangeEventListener(HAND, EXILE)
}