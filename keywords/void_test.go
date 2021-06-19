package keywords

import (
	"testing"
)

func TestVoid(t *testing.T) {
	void := void{"TARGET"}
	err := void.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
