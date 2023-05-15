package keywords

import (
	"testing"
)

func TestLaneswap(t *testing.T) {
	laneswap := laneswap{"ALL"}
	err := laneswap.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
