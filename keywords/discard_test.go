package keywords

import (
	"testing"
)

func TestDiscard(t *testing.T) {
	discard := discard{}
	err := discard.Validate()
	if err != nil {
		t.Error(err)
	}
}
