package keywords

import (
	"testing"
)

func TestHarm(t *testing.T) {
	harm := harm{3}
	err := harm.Validate()
	if err != nil {
		t.Error(err)
	}
}
