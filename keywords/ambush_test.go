package keywords

import (
	"testing"
)

func TestAmbush(t *testing.T) {
	ambush := bounce{"THIS"}
	err := ambush.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
