package keywords

import (
	"testing"
)

func TestRecoverEntity(t *testing.T) {
	recoverEntity := recoverEntity{}
	err := recoverEntity.Validate()
	if err != nil {
		t.Error(err)
	}
}
