package keywords

import (
	"testing"
)

func TestAmbush(t *testing.T) {
	ambush := bounce{"THIS"}
	err := ambush.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
