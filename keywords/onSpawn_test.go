package keywords

import (
	"testing"
)

func TestOnSpawn(t *testing.T) {
	onSpawn := onSpawn{}
	err := onSpawn.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
