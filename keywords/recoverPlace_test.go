package keywords

import (
	"testing"
)

func TestRecoverPlace(t *testing.T) {
	recoverPlace := recoverPlace{}
	err := recoverPlace.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
