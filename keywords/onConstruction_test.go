package keywords

import (
	"testing"
)

func TestOnConstruction(t *testing.T) {
	onConstruction := onConstruction{nil}
	err := onConstruction.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
