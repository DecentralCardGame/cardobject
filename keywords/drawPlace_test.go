package keywords

import (
	"testing"
)

func TestDrawPlace(t *testing.T) {
	drawPlace := drawPlace{nil}
	err := drawPlace.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
