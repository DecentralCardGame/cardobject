package keywords

import (
	"testing"
)

func TestBattlecry(t *testing.T) {
	battlecry := battlecry{}
	err := battlecry.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
