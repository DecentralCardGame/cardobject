package keywords

import (
	"testing"
)

func TestGrow(t *testing.T) {
	grow := grow{4}
	err := grow.Validate()
	if err != nil {
		t.Error(err)
	}
}
