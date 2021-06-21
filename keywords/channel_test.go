package keywords

import (
	"testing"
)

func TestChannel(t *testing.T) {
	channel := channel{nil}
	err := channel.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
