package keywords

import (
	"testing"
)

func TestArmor(t *testing.T) {
	armor := armor{3}
	err := armor.Validate()
	if err != nil {
		t.Error(err)
	}
}
