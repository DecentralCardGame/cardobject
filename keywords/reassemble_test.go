package keywords

import (
	"testing"
)

func TestReassemble(t *testing.T) {
	reassemble := reassemble{}
	err := reassemble.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
