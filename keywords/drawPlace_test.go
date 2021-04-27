package keywords

import (
	"testing"
)

func TestDrawPlace(t *testing.T) {
	drawPlace := drawPlace{nil}
	err := drawPlace.Validate()
	if err != nil {
		t.Error(err)
	}
}
