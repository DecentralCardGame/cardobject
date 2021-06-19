package keywords

import (
	"testing"
)

func TestPay(t *testing.T) {
	pay := pay{3, nil}
	err := pay.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
