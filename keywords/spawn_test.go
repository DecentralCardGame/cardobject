package keywords

import (
	"testing"

	"github.com/DecentralCardGame/cardobject/cardobject"
)

func TestSpawn(t *testing.T) {
	intVariable := cardobject.IntVariableName("X")
	intValue := cardobject.IntValue{SimpleIntValue: nil, IntVariable: &intVariable}
	spawn := spawn{"2/2 Bot", intValue}
	err := spawn.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
