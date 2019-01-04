package cardobject

import "testing"
import "reflect"

func TestZoneChange(t *testing.T) {
	testedInterface := "ZoneChange"
	x := createZoneChange()
	if(!reflect.DeepEqual(x.GetSelection(), createSafeSelector())) {
		t.Errorf(testedInterface + " doesn't return CardSelection")
	}
	if(x.GetDestination() != EXILE) {
		t.Errorf(testedInterface + " doesn't return Destination")
	}
	success(testedInterface)
}

func createZoneChange() ZoneChange {
	return NewZoneChange(createSafeSelector(), EXILE)
}