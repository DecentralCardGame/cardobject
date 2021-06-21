package keywords

import (
	"testing"
)

func TestPay(t *testing.T) {
	pay := pay{3, nil}
	err := pay.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
