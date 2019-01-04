package cardobject

import "testing"

const testIntConst = 5
const testStringConst = "test"

func TestIntInserter(t *testing.T) {
	testedInterface := "IntInserter"
	x := createIntConst()
	if(x.GetIntVal() != testIntConst || !x.IsIntRetriever()){
		t.Errorf("IntInserterConst doesn't implement " + testedInterface)
	}
	if(ATTACK.GetIntPropertyId() != ATTACK || !ATTACK.IsIntRetriever()) {
		t.Errorf("IntPropId doesn't implement " + testedInterface)
	}
	success(testedInterface)
}

func TestStringInserter(t *testing.T) {
	testedInterface := "StringInserter"
	x := createStringConst()
	if(x.GetStringVal() != testStringConst || !x.IsStringRetriever()){
		t.Errorf("StringInserterConst doesn't implement " + testedInterface)
	}
	if(TAG.GetStringPropertyId() != TAG || !TAG.IsStringRetriever()) {
		t.Errorf("StringPropId doesn't implement " + testedInterface)
	}
	success(testedInterface)
}

func createIntConst() IntInserterConst {
	return NewIntConst(testIntConst)
}

func createStringConst() StringInserterConst {
	return NewStringConst(testStringConst)
}