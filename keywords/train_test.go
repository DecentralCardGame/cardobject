package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestTrain(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	train := train{"RANDOM", intValue}
	err := train.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
