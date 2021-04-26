package keywords

import (
	"testing"
)

func TestKill(t *testing.T) {
	kill := kill{}
	err := kill.Validate()
	if err != nil {
		t.Error(err)
	}
}
