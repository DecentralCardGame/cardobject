package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestDice(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	dice := dice{intValue}
	err := dice.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
