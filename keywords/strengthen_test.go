package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestStrengthen(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	strengthen := strengthen{"ALL", intValue}
	err := strengthen.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
