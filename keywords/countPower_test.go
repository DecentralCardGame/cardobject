package keywords

import (
	"testing"
)

func TestAmbush(t *testing.T) {
	countPower := countPower{"THIS"}
	err := countPower.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
