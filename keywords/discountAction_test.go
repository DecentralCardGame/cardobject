package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestDiscountAction(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	discountAction := discountAction{intValue}
	err := discountAction.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
