package keywords

import (
	"testing"
)

func TestArrival(t *testing.T) {
	arrival := arrival{}
	err := arrival.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
