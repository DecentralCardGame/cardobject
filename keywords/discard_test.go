package keywords

import (
	"testing"
)

func TestDiscard(t *testing.T) {
	discard := discard{}
	err := discard.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
