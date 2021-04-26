package keywords

import (
	"testing"
)

func TestTribute(t *testing.T) {
	tribute := tribute{}
	err := tribute.Validate()
	if err != nil {
		t.Error(err)
	}
}
