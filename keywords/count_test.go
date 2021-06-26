package keywords

import (
	"testing"
)

func TestCount(t *testing.T) {
	count := count{"DWARF"}
	err := count.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
