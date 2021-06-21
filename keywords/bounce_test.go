package keywords

import (
	"testing"
)

func TestBounce(t *testing.T) {
	bounce := bounce{"ALL"}
	err := bounce.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
