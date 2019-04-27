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

func TestName(t *testing.T) {
	x := createName()
	if (x.CardStringPropertyId() != NAME) {
		t.Errorf("NameId was incorrect")
	}
	if (x.GetStringVal() != testPropIdString) {
		t.Errorf("NameVal was incorrect")
	}
	success("Name")
}

func TestSpeedModifier(t *testing.T) {
	x := createSpeedModifier()
	if (x.CardIntPropertyId() != SPEEDMODIFIER) {
		t.Errorf("SpeedModifierId was incorrect")
	}
	if (x.GetIntVal() != testPropIdInt) {
		t.Errorf("SpeedModifierVal was incorrect")
	}
	success("SpeedModifier")
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

func TestUniqueName(t *testing.T) {
	x := createUniqueName()
	if (x.CardStringPropertyId() != UNIQUENAME) {
		t.Errorf("UniqueNameId was incorrect")
	}
	if (x.GetStringVal() != (testPropIdString + ", the " + testPropIdString)) {
		t.Errorf("UniqueNameVal was incorrect")
	}
	success("UniqueName")
}

func createAttack() Attack {
	return NewAttack(testPropIdInt)
}

func createCostSum() CostSum {
	return NewCostSum(testPropIdInt)
}

func createHealth() Health {
	return NewHealth(testPropIdInt)
}

func createName() Name {
	return NewName(testPropIdString)
}

func createSpeedModifier() SpeedModifier {
	return NewSpeedModifier(testPropIdInt)
}

func createTag() Tag {
	return NewTag(testPropIdString)
}

func createText() Text {
	return NewText(testPropIdString)
}

func createUniqueName() UniqueName {
	return NewUniqueName(testPropIdString, testPropIdString)
}