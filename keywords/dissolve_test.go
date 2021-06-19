package keywords

import (
	"testing"
)

func TestDissolve(t *testing.T) {
	dissolve := dissolve{3, nil}
	err := dissolve.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
