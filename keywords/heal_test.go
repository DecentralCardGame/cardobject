package keywords

import (
	"testing"
)

func TestHeal(t *testing.T) {
	heal := heal{}
	err := heal.Validate()
	if err != nil {
		t.Error(err)
	}
}
