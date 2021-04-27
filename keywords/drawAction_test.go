package keywords

import (
	"testing"
)

func TestDrawAction(t *testing.T) {
	drawAction := drawAction{nil}
	err := drawAction.Validate()
	if err != nil {
		t.Error(err)
	}
}
