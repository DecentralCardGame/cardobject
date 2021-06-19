package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestArmor(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	armor := armor{intValue}
	err := armor.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
