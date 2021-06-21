package keywords

import (
	"testing"
)

func TestDiscardPay(t *testing.T) {
	discardPay := discardPay{3, nil}
	err := discardPay.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
