package keywords

import (
	"testing"
)

func TestChannel(t *testing.T) {
	channel := channel{nil}
	err := channel.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
