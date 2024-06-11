package keywords

import (
	"testing"
)

func TestCountDust(t *testing.T) {
	countDust := countDust{"ACTION"}
	err := countDust.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
