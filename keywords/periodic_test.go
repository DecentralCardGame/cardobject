package keywords

import (
	"testing"
)

func TestPeriodic(t *testing.T) {
	periodic := periodic{}
	err := periodic.Validate()
	if err != nil {
		t.Error(err)
	}
}
