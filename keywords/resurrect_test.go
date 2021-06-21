package keywords

import (
	"testing"
)

func TestResurrect(t *testing.T) {
	resurrect := resurrect{}
	err := resurrect.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
