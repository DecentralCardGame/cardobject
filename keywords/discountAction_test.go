package keywords

import "testing"

func TestDiscountAction(t *testing.T) {
	discountAction := discountAction{}
	err := discountAction.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
