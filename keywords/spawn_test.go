package keywords

import (
	"testing"
)

func TestSpawn(t *testing.T) {
	spawn := spawn{"2/2 Bot", 5}
	err := spawn.Validate()
	if err != nil {
		t.Error(err)
	}
}
