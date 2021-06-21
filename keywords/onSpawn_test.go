package keywords

import (
	"testing"
)

func TestOnSpawn(t *testing.T) {
	onSpawn := onSpawn{}
	err := onSpawn.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
