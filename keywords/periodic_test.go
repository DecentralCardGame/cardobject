package keywords

import (
	"testing"
)

func TestPeriodic(t *testing.T) {
	periodic := periodic{}
	err := periodic.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
