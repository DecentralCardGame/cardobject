package keywords

import (
	"testing"
)

func TestArm(t *testing.T) {
	arm := arm{"THIS", 3}
	err := arm.Validate()
	if err != nil {
		t.Error(err)
	}
}
