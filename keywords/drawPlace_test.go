package keywords

import (
	"testing"
)

func TestDrawPlace(t *testing.T) {
	drawPlace := drawPlace{nil}
	err := drawPlace.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
