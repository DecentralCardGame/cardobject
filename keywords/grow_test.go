package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestGrow(t *testing.T) {
	simpleIntValue := cardobject.SimpleIntValue(3)
	intValue := cardobject.IntValue{SimpleIntValue: &simpleIntValue, IntVariable: nil}
	grow := grow{intValue}
	err := grow.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}

func TestGrowDescription(t *testing.T) {
	const want = "Grow your HQ, 4 Growth increases your HQ's mana production each turn by 1 permanently."
	if got := (grow{}).Description(); got != want {
		t.Fatalf("Grow description = %q, want %q", got, want)
	}
}
