package keywords

import (
	"testing"
)

func TestAnthem(t *testing.T) {
	anthem := anthem{"DWARF"}
	err := anthem.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
