package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestProduce(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	produce := produce{intValue}
	err := produce.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
