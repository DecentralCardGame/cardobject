package keywords

import (
	"testing"
)

func TestRavage(t *testing.T) {
	ravage := ravage{"RANDOM", 3}
	err := ravage.Validate()
	if err != nil {
		t.Error(err)
	}
}
