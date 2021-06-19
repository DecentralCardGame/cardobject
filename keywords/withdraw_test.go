package keywords

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	withdraw := withdraw{"ALL"}
	err := withdraw.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
