package keywords

import (
	"testing"
)

func TestRecoverAction(t *testing.T) {
	recoverAction := recoverAction{}
	err := recoverAction.Validate()
	if err != nil {
		t.Error(err)
	}
}
