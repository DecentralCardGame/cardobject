package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestSpawn(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	token := token{nil, nil, nil, nil, &recruit{}, nil}
	spawn := spawn{token, intValue}
	err := spawn.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
