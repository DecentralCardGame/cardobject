package keywords

import (
	"testing"
)

func TestSelfBurn(t *testing.T) {
	selfBurn := selfBurn{3}
	err := selfBurn.Validate()
	if err != nil {
		t.Error(err)
	}
}
