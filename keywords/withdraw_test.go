package keywords

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	withdraw := withdraw{"ALL"}
	err := withdraw.Validate()
	if err != nil {
		t.Error(err)
	}
}
