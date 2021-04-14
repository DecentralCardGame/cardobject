package cardobject

import (
	"io/ioutil"
	"testing"
)

func TestValidateAction(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/actionTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := Unmarshal(dat)

	if err != nil {
		t.Error(err)
	}
}

func TestValidateEntity(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/entityTest1.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := Unmarshal(dat)

	if err != nil {
		t.Error(err)
	}
}

func TestValidatePlace(t *testing.T) {
	dat, readErr := ioutil.ReadFile("testJsons/place1Test.json")
	if readErr != nil {
		t.Error(readErr)
	}

	_, err := Unmarshal(dat)

	if err != nil {
		t.Error(err)
	}
}
