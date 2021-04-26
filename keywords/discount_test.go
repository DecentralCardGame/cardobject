package keywords

import (
	"testing"
)

func TestDiscount(t *testing.T) {
	discount := discount{3}
	err := discount.Validate()
	if err != nil {
		t.Error(err)
	}
}
