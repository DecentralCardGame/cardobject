package keywords

import (
	"testing"
)

func TestAnthem(t *testing.T) {
	anthem := anthem{"DWARF"}
	err := anthem.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
