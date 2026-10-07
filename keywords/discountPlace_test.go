package keywords

import "testing"

func TestDiscountPlace(t *testing.T) {
	discountPlace := discountPlace{}
	err := discountPlace.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
