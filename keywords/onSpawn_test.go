package keywords

import (
	"testing"
)

func TestOnSpawn(t *testing.T) {
	onSpawn := onSpawn{}
	err := onSpawn.Validate()
	if err != nil {
		t.Error(err)
	}
}
