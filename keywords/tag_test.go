package keywords

import (
	"testing"
)

func TestTag(t *testing.T) {
	tag := tag{"TACTIC", nil}
	err := tag.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
