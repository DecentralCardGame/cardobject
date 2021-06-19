package keywords

import (
	"testing"
)

func TestDrawEntity(t *testing.T) {
	drawEntity := drawEntity{nil}
	err := drawEntity.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
