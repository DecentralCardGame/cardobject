package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestArmor(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	armor := armor{"THIS", intValue}
	err := armor.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
