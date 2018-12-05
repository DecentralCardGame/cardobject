package cardobject

import "testing"

const testIntConst = 5
const testStringConst = "test"

func TestIntInserter(t *testing.T) {
	x := createIntConst()
	if(x.GetIntVal() != testIntConst || !x.IsIntRetriever()){
		t.Errorf("IntInserterConst doesn't implement Inserter")
	}
	if(ATTACK.GetIntPropertyId() != ATTACK || !ATTACK.IsIntRetriever()) {
		t.Errorf("IntPropId doesn't implement Inserter")
	}
	success("IntInserter")
}

func TestStringInserter(t *testing.T) {
	x := createStringConst()
	if(x.GetStringVal() != testStringConst || !x.IsStringRetriever()){
		t.Errorf("StringInserterConst doesn't implement Inserter")
	}
	if(TAG.GetStringPropertyId() != TAG || !TAG.IsStringRetriever()) {
		t.Errorf("StringPropId doesn't implement Inserter")
	}
	success("StringInserter")
}

func createIntConst() IntInserterConst {
	return NewIntConst(testIntConst)
}

func createStringConst() StringInserterConst {
	return NewStringConst(testStringConst)
}