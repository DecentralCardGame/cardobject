package keywords

import (
	"testing"
)

func TestCountPower(t *testing.T) {
	countForces := countForces{3}
	err := countForces.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
