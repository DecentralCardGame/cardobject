package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestFortify(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	fortify := fortify{intValue}
	err := fortify.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
