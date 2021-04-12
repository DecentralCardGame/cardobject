package keywords

import (
	"testing"
)

func TestAnthem(t *testing.T) {
	anthem := anthem{"DWARF"}
	err := anthem.Validate()
	if err != nil {
		t.Error(err)
	}
	effect := anthem.Resolve()
	err = effect.Validate()
	if err != nil {
		t.Error(err)
	}
}
