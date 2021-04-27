package keywords

import (
	"testing"
)

func TestRegenerate(t *testing.T) {
	regenerate := regenerate{3}
	err := regenerate.Validate()
	if err != nil {
		t.Error(err)
	}
}
