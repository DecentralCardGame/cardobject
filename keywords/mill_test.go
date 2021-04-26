package keywords

import (
	"testing"
)

func TestMill(t *testing.T) {
	mill := mill{6, nil}
	err := mill.Validate()
	if err != nil {
		t.Error(err)
	}
}
