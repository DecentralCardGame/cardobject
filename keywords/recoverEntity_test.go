package keywords

import (
	"testing"
)

func TestRecoverEntity(t *testing.T) {
	recoverEntity := recoverEntity{}
	err := recoverEntity.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
