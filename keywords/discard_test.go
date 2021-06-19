package keywords

import (
	"testing"
)

func TestDiscard(t *testing.T) {
	discard := discard{}
	err := discard.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
