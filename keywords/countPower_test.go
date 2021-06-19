package keywords

import (
	"testing"
)

func TestCountPower(t *testing.T) {
	countPower := countPower{3}
	err := countPower.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
