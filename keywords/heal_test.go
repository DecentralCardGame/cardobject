package keywords

import (
	"testing"
)

func TestHeal(t *testing.T) {
	heal := heal{"TARGET"}
	err := heal.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
