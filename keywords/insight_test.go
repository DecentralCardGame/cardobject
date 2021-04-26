package keywords

import (
	"testing"
)

func TestInsight(t *testing.T) {
	insight := insight{4}
	err := insight.Validate()
	if err != nil {
		t.Error(err)
	}
}
