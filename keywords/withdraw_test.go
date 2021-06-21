package keywords

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	withdraw := withdraw{"ALL"}
	err := withdraw.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
