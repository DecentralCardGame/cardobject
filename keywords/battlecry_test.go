package keywords

import (
	"testing"
)

func TestBattlecry(t *testing.T) {
	battlecry := battlecry{}
	err := battlecry.Validate()
	if err != nil {
		t.Error(err)
	}
}
