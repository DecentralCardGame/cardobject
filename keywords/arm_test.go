package keywords

import (
	"testing"
)

func TestArm(t *testing.T) {
	arm := arm{3}
	err := arm.Validate()
	if err != nil {
		t.Error(err)
	}
	effect := arm.Resolve()
	err = effect.Validate()
	if err != nil {
		t.Error(err)
	}
}
