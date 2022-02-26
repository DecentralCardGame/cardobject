package keywords

import (
	"testing"
)

func TestTrample(t *testing.T) {
	trample := trample{}
	err := trample.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
