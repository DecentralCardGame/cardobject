package keywords

import "testing"

func TestDiscountEntity(t *testing.T) {
	discountEntity := discountEntity{}
	err := discountEntity.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
