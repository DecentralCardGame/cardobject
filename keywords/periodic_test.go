package keywords

import (
	"testing"
)

func TestPeriodic(t *testing.T) {
	periodic := periodic{}
	err := periodic.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
