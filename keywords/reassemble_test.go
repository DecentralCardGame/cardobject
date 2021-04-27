package keywords

import (
	"testing"
)

func TestReassemble(t *testing.T) {
	reassemble := reassemble{}
	err := reassemble.Validate()
	if err != nil {
		t.Error(err)
	}
}
