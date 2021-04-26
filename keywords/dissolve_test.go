package keywords

import (
	"testing"
)

func TestDissolve(t *testing.T) {
	dissolve := dissolve{3, nil}
	err := dissolve.Validate()
	if err != nil {
		t.Error(err)
	}
}
