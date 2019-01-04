package cardobject

import "testing"

const TestIntCondition = 5
const testIntCondition = 5
const testStringCondition = "test"

func TestCardIntCondition(t *testing.T) {
	testedInterface := "CardIntCondition"
	x := createCardIntCondition()
	if(x.GetComparator() != GREATER) {
		t.Errorf(testedInterface + " doesn't return Comparator")
	}
	if(x.GetCompVal() != testIntCondition) {
		t.Errorf(testedInterface + " doesn't return CompVal")
	}
	success(testedInterface)
}

func TestCardStringCondition(t *testing.T) {
	testedInterface := "CardStringCondition"
	x := createCardStringCondition()
	if(x.GetComparator() != EQUAL) {
		t.Errorf(testedInterface + " doesn't return Comparator")
	}
	if(x.GetCompVal() != testStringCondition) {
		t.Errorf(testedInterface + " doesn't return CompVal")
	}
	success(testedInterface)
}

func TestPlayerIntCondition(t *testing.T) {
	testedInterface := "PlayerIntCondition"
	x := createPlayerIntCondition()
	if(x.GetComparator() != GREATER) {
		t.Errorf(testedInterface + " doesn't return Comparator")
	}
	if(x.GetCompVal() != testIntCondition) {
		t.Errorf(testedInterface + " doesn't return CompVal")
	}
	success(testedInterface)
}

func createCardIntCondition() CardIntCondition {
	return NewCardIntCondition(GREATER, TestIntCondition, ATTACK)
}

func createCardStringCondition() CardStringCondition {
	return NewCardStringCondition(testStringCondition, TAG)
}

func createPlayerIntCondition() PlayerIntCondition {
	return NewPlayerIntCondition(GREATER, testIntCondition, HANDSIZE)
}

