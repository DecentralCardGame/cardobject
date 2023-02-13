package keywords

import (
	"testing"
)

func TestEnrage(t *testing.T) {
	enrage := enrage{"TARGET"}
	err := enrage.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
