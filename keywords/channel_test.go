package keywords

import (
	"testing"
)

func TestChannel(t *testing.T) {
	channel := channel{nil}
	err := channel.Validate()
	if err != nil {
		t.Error(err)
	}
}
