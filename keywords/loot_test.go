package keywords

import (
	"testing"
)

func TestLoot(t *testing.T) {
	loot := loot{nil}
	err := loot.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
