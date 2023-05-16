package keywords

import (
	"testing"
)

func TestChoose(t *testing.T) {
	choose := choose{"ENTITY", "HAND"}
	err := choose.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
