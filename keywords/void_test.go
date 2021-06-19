package keywords

import (
	"testing"
)

func TestVoid(t *testing.T) {
	void := void{"TARGET"}
	err := void.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
