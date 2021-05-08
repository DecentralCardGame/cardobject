package keywords

import (
	"testing"
)

func TestStrengthen(t *testing.T) {
	strengthen := strengthen{"ALL", 3}
	err := strengthen.Validate()
	if err != nil {
		t.Error(err)
	}
}
