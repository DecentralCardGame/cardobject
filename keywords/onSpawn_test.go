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
	ability := onSpawn.Resolve()
	err = ability.Validate()
	if err != nil {
		t.Error(err)
	}
}
