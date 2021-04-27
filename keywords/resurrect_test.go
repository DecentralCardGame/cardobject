package keywords

import (
	"testing"
)

func TestResurrect(t *testing.T) {
	resurrect := resurrect{}
	err := resurrect.Validate()
	if err != nil {
		t.Error(err)
	}
}
