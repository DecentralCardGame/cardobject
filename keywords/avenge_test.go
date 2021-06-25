package keywords

import (
	"testing"
)

func TestAvenge(t *testing.T) {
	avenge := avenge{}
	err := avenge.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
