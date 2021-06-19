package keywords

import (
	"testing"
)

func TestKill(t *testing.T) {
	kill := kill{"TARGET"}
	err := kill.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
