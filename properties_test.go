package cardobject

import "testing"

const testInt = 5
const testString = "test"

func TestAttack(t *testing.T) {
	x := createAttack(testInt)
	if (x.CardIntPropertyId() != ATTACK) {
		t.Errorf("AttackId was incorrect")
	}
	if (x.GetIntVal() != testInt) {
		t.Errorf("AttackVal was incorrect")
	}
	success("Attack")
}

func TestCostSum(t *testing.T) {
	x := createCostSum(testInt)
	if (x.CardIntPropertyId() != COSTSUM) {
		t.Errorf("CostSumId was incorrect")
	}
	if (x.GetIntVal() != testInt) {
		t.Errorf("CostSumVal was incorrect")
	}
	success("CostSum")
}

func TestHealth(t *testing.T) {
	x := createHealth(testInt)
	if (x.CardIntPropertyId() != HEALTH) {
		t.Errorf("HealthId was incorrect")
	}
	if (x.GetIntVal() != testInt) {
		t.Errorf("HealthVal was incorrect")
	}
	success("Health")
}

func TestSpeedmodifier(t *testing.T) {
	x := createSpeedmodifier(testInt)
	if (x.CardIntPropertyId() != SPEEDMODIFIER) {
		t.Errorf("SpeedmodifierId was incorrect")
	}
	if (x.GetIntVal() != testInt) {
		t.Errorf("SpeedmodifierVal was incorrect")
	}
	success("Speedmodifier")
}

func TestTag(t *testing.T) {
	x := createTag(testString)
	if (x.CardStringPropertyId() != TAG) {
		t.Errorf("TagId was incorrect")
	}
	if (x.GetStringVal() != testString) {
		t.Errorf("TagVal was incorrect")
	}
	success("Tag")
}

func TestText(t *testing.T) {
	x := createText(testString)
	if (x.CardStringPropertyId() != TEXT) {
		t.Errorf("TextId was incorrect")
	}
	if (x.GetStringVal() != testString) {
		t.Errorf("TextVal was incorrect")
	}
	success("Text")
}

func createAttack(i int) *Attack {
	return NewAttack(i)
}

func createCostSum(i int) *CostSum {
	return NewCostSum(i)
}

func createHealth(i int) *Health {
	return NewHealth(i)
}

func createSpeedmodifier(i int) *Speedmodifier {
	return NewSpeedmodifier(i)
}

func createTag(s string) *Tag {
	return NewTag(s)
}

func createText(s string) *Text {
	return NewText(s)
}