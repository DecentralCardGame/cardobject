package keywords

import (
	"testing"
)

func TestDismantle(t *testing.T) {
	dismantle := dismantle{nil}
	err := dismantle.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
