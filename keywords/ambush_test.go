package keywords

import (
	"testing"
)

func TestAmbush(t *testing.T) {
	ambush := bounce{"THIS"}
	err := ambush.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
