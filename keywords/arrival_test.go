package keywords

import (
	"testing"
)

func TestArrival(t *testing.T) {
	arrival := arrival{}
	err := arrival.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
