package keywords

import (
	"testing"
)

func TestHeal(t *testing.T) {
	heal := heal{"TARGET"}
	err := heal.Validate()
	if err != nil {
		t.Error(err)
	}
}
