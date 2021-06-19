package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestRavage(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	ravage := ravage{"RANDOM", intValue}
	err := ravage.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
