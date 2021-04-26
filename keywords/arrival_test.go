package keywords

import (
	"testing"
)

func TestArrival(t *testing.T) {
	arrival := arrival{}
	err := arrival.Validate()
	if err != nil {
		t.Error(err)
	}
}
