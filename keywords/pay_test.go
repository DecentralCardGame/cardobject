package keywords

import (
	"testing"
)

func TestPay(t *testing.T) {
	pay := pay{3, nil}
	err := pay.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
