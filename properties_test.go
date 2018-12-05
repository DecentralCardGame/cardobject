package cardobject

import "testing"

const testPropIdInt = 5
const testPropIdString = "test"

func TestAttack(t *testing.T) {
	x := createAttack()
	if (x.CardIntPropertyId() != ATTACK) {
		t.Errorf("AttackId was incorrect")
	}
	if (x.GetIntVal() != testPropIdInt) {
		t.Errorf("AttackVal was incorrect")
	}
	success("Attack")
}

func TestCostSum(t *testing.T) {
	x := createCostSum()
	if (x.CardIntPropertyId() != COSTSUM) {
		t.Errorf("CostSumId was incorrect")
	}
	if (x.GetIntVal() != testPropIdInt) {
		t.Errorf("CostSumVal was incorrect")
	}
	success("CostSum")
}

func TestHealth(t *testing.T) {
	x := createHealth()
	if (x.CardIntPropertyId() != HEALTH) {
		t.Errorf("HealthId was incorrect")
	}
	if (x.GetIntVal() != testPropIdInt) {
		t.Errorf("HealthVal was incorrect")
	}
	success("Health")
}

func TestSpeedmodifier(t *testing.T) {
	x := createSpeedmodifier()
	if (x.CardIntPropertyId() != SPEEDMODIFIER) {
		t.Errorf("SpeedmodifierId was incorrect")
	}
	if (x.GetIntVal() != testPropIdInt) {
		t.Errorf("SpeedmodifierVal was incorrect")
	}
	success("Speedmodifier")
}

func TestTag(t *testing.T) {
	x := createTag()
	if (x.CardStringPropertyId() != TAG) {
		t.Errorf("TagId was incorrect")
	}
	if (x.GetStringVal() != testPropIdString) {
		t.Errorf("TagVal was incorrect")
	}
	success("Tag")
}

func TestText(t *testing.T) {
	x := createText()
	if (x.CardStringPropertyId() != TEXT) {
		t.Errorf("TextId was incorrect")
	}
	if (x.GetStringVal() != testPropIdString) {
		t.Errorf("TextVal was incorrect")
	}
	success("Text")
}

func createAttack() *Attack {
	return NewAttack(testPropIdInt)
}

func createCostSum() *CostSum {
	return NewCostSum(testPropIdInt)
}

func createHealth() *Health {
	return NewHealth(testPropIdInt)
}

func createSpeedmodifier() *Speedmodifier {
	return NewSpeedmodifier(testPropIdInt)
}

func createTag() *Tag {
	return NewTag(testPropIdString)
}

func createText() *Text {
	return NewText(testPropIdString)
}