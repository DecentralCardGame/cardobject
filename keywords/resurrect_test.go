package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestResurrect(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	resurrect := resurrect{intValue}
	err := resurrect.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
