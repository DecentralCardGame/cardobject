package keywords

import (
	"testing"
)

func TestReassemble(t *testing.T) {
	reassemble := reassemble{}
	err := reassemble.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
