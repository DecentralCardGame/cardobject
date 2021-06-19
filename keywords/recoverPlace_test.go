package keywords

import (
	"testing"
)

func TestRecoverPlace(t *testing.T) {
	recoverPlace := recoverPlace{}
	err := recoverPlace.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
