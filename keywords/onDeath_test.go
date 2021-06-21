package keywords

import (
	"testing"
)

func TestOnDeath(t *testing.T) {
	onDeath := onDeath{}
	err := onDeath.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
