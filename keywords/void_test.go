package keywords

import (
	"testing"
)

func TestVoid(t *testing.T) {
	void := void{"TARGET"}
	err := void.Validate()
	if err != nil {
		t.Error(err)
	}
}
