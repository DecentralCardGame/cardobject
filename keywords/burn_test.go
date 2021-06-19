package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestBurn(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	burn := burn{intValue}
	err := burn.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
