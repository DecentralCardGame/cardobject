package keywords

import (
	"testing"
)

func TestCount(t *testing.T) {
	count := count{"TACTIC"}
	err := count.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
