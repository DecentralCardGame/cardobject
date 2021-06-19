package keywords

import (
	"testing"
)

func TestDrawAction(t *testing.T) {
	drawAction := drawAction{nil}
	err := drawAction.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
