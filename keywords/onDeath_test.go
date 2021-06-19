package keywords

import (
	"testing"
)

func TestOnDeath(t *testing.T) {
	onDeath := onDeath{}
	err := onDeath.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
