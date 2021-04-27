package keywords

import (
	"testing"
)

func TestDrawEntity(t *testing.T) {
	drawEntity := drawEntity{nil}
	err := drawEntity.Validate()
	if err != nil {
		t.Error(err)
	}
}
