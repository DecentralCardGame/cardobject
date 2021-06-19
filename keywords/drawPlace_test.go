package keywords

import (
	"testing"
)

func TestDrawPlace(t *testing.T) {
	drawPlace := drawPlace{nil}
	err := drawPlace.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
