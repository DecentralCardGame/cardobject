package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestPay(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	pay := pay{intValue, nil}
	err := pay.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
