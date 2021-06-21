package keywords

import (
	"testing"
)

func TestKill(t *testing.T) {
	kill := kill{"TARGET"}
	err := kill.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
