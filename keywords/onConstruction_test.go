package keywords

import (
	"testing"
)

func TestOnConstruction(t *testing.T) {
	onConstruction := onConstruction{}
	err := onConstruction.Validate()
	if err != nil {
		t.Error(err)
	}
	ability := onConstruction.Resolve()
	err = ability.Validate()
	if err != nil {
		t.Error(err)
	}
}
