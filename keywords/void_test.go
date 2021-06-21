package keywords

import (
	"testing"
)

func TestVoid(t *testing.T) {
	void := void{"TARGET"}
	err := void.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
