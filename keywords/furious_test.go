package keywords

import (
	"testing"
)

func TestFurious(t *testing.T) {
	furious := furious{nil}
	err := furious.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
