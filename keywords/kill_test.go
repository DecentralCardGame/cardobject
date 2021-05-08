package keywords

import (
	"testing"
)

func TestKill(t *testing.T) {
	kill := kill{"TARGET"}
	err := kill.Validate()
	if err != nil {
		t.Error(err)
	}
}
