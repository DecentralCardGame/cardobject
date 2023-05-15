package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestDiscountPlace(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	discountPlace := discountPlace{intValue}
	err := discountPlace.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
