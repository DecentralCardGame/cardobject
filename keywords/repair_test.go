package keywords

import (
	"testing"
)

func TestRepair(t *testing.T) {
	repair := repair{3}
	err := repair.Validate()
	if err != nil {
		t.Error(err)
	}
}
