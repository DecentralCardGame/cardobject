package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestDiscountEntity(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	discountEntity := discountEntity{intValue}
	err := discountEntity.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
