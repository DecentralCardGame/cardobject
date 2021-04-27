package keywords

import (
	"testing"
)

func TestLoot(t *testing.T) {
	loot := loot{nil}
	err := loot.Validate()
	if err != nil {
		t.Error(err)
	}
}
