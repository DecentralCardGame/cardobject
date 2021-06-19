package keywords

import (
	"testing"
)

func TestSilence(t *testing.T) {
	silence := silence{"ALL", "YOU"}
	err := silence.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
