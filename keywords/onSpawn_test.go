package keywords

import (
	"testing"
)

func TestOnSpawn(t *testing.T) {
	onSpawn := onSpawn{}
	err := onSpawn.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
