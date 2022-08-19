package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestReassemble(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	reassemble := reassemble{intValue}
	err := reassemble.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
