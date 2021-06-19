package keywords

import (
	"testing"
)

func TestOnConstruction(t *testing.T) {
	onConstruction := onConstruction{nil}
	err := onConstruction.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
