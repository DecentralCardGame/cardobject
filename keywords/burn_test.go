package keywords

import (
	"testing"
)

func TestBurn(t *testing.T) {
	burn := burn{3}
	err := burn.Validate()
	if err != nil {
		t.Error(err)
	}
}
