package keywords

import (
	"testing"
)

func TestHasten(t *testing.T) {
	hasten := hasten{"ALL"}
	err := hasten.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
