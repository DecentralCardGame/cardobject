package keywords

import (
	"testing"
)

func TestDrawEntity(t *testing.T) {
	drawEntity := drawEntity{nil}
	err := drawEntity.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
