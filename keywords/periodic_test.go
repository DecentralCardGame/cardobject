package keywords

import (
	"testing"
)

func TestPeriodic(t *testing.T) {
	periodic := periodic{}
	err := periodic.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
