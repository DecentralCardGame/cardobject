package keywords

import (
	"testing"
)

func TestBounce(t *testing.T) {
	bounce := bounce{"ALL"}
	err := bounce.Validate()
	if err != nil {
		t.Error(err)
	}
}
