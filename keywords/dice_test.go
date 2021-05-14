package keywords

import (
	"testing"
)

func TestDice(t *testing.T) {
	dice := dice{3}
	err := dice.Validate()
	if err != nil {
		t.Error(err)
	}
}
