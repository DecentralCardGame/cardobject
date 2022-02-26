package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestGrind(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	grind := grind{"ALL", intValue}
	err := grind.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
