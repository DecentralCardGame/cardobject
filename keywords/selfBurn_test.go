package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestSelfBurn(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	selfBurn := selfBurn{intValue}
	err := selfBurn.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
