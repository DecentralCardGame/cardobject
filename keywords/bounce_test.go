package keywords

import (
	"testing"
)

func TestBounce(t *testing.T) {
	bounce := bounce{"ALL"}
	err := bounce.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
