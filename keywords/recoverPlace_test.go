package keywords

import (
	"testing"
)

func TestRecoverPlace(t *testing.T) {
	recoverPlace := recoverPlace{}
	err := recoverPlace.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
