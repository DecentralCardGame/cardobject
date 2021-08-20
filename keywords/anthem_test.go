package keywords

import (
	"testing"
)

func TestAnthem(t *testing.T) {
	anthem := anthem{"TACTIC"}
	err := anthem.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
