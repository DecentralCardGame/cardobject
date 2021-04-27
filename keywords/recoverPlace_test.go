package keywords

import (
	"testing"
)

func TestRecoverPlace(t *testing.T) {
	recoverPlace := recoverPlace{}
	err := recoverPlace.Validate()
	if err != nil {
		t.Error(err)
	}
}
