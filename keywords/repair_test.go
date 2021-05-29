package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestRepair(t *testing.T) {
	intVaraiable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVaraiable}
	repair := repair{"ALL", intValue}
	err := repair.Validate()
	if err != nil {
		t.Error(err)
	}
}
