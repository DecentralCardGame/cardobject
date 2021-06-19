package keywords

import (
	"testing"
)

func TestOnDeath(t *testing.T) {
	onDeath := onDeath{}
	err := onDeath.ValidateType(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
