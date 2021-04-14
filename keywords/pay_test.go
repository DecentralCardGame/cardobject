package keywords

import (
	"testing"
)

func TestPay(t *testing.T) {
	pay := pay{3, nil}
	err := pay.Validate()
	if err != nil {
		t.Error(err)
	}
	ability := pay.Resolve()
	err = ability.Validate()
	if err != nil {
		t.Error(err)
	}
}
