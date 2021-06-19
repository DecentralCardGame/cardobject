package keywords

import (
	"testing"
)

func TestRecoverEntity(t *testing.T) {
	recoverEntity := recoverEntity{}
	err := recoverEntity.Validate(emptyTestCard())
	if err != nil {
		t.Error(err)
	}
}
