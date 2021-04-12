package keywords

import (
	"testing"
)

func TestProduce(t *testing.T) {
	produce := produce{4}
	err := produce.Validate()
	if err != nil {
		t.Error(err)
	}
	effect := produce.Resolve()
	err = effect.Validate()
	if err != nil {
		t.Error(err)
	}
}