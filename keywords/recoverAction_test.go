package keywords

import (
	"testing"
)

func TestRecoverAction(t *testing.T) {
	recoverAction := recoverAction{}
	err := recoverAction.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
