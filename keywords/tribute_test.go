package keywords

import (
	"testing"
)

func TestTribute(t *testing.T) {
	tribute := tribute{}
	err := tribute.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
