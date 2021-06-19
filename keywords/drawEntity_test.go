package keywords

import (
	"testing"
)

func TestDrawEntity(t *testing.T) {
	drawEntity := drawEntity{nil}
	err := drawEntity.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
