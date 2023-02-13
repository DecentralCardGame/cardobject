package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestMill(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	mill := mill{intValue, "YOU"}
	err := mill.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
