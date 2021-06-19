package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestGrow(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	grow := grow{intValue}
	err := grow.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
