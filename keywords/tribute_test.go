package keywords

import (
	"testing"
)

func TestTribute(t *testing.T) {
	tribute := tribute{"RANDOM", nil}
	err := tribute.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
