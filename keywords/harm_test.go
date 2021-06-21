package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestHarm(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	harm := harm{"RANDOM", intValue}
	err := harm.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
