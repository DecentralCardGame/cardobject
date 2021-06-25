package keywords

import (
	"testing"
)

func TestSacrifice(t *testing.T) {
	sacrifice := sacrifice{"TARGET"}
	err := sacrifice.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
