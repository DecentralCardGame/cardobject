package cardobject

import "testing"
import "reflect"
import "fmt"

func TestAction(t *testing.T) {
	testedInterface := "Action"
	x := createAction()
	if(!reflect.DeepEqual(x.GetName(), createName())) {
		t.Errorf(testedInterface + " doesn't return Name")
	}
	if(!reflect.DeepEqual(x.GetCost(), createCost())) {
		t.Errorf(testedInterface + " doesn't return Cost")
	}
	if(!reflect.DeepEqual(x.GetCastSpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return SpeedModifier")
	}
	if(!reflect.DeepEqual(x.GetTags(), []Tag{createTag()})) {
		t.Errorf(testedInterface + " doesn't return Tags")
	}
	if(!reflect.DeepEqual(x.GetText(), createText())) {
		t.Errorf(testedInterface + " doesn't return Text")
	}
	if(!reflect.DeepEqual(x.GetEffect(), createEffect())) {
		t.Errorf(testedInterface + " doesn't return Effect")
	}
	success(testedInterface)
}

func TestEntity(t *testing.T) {
	testedInterface := "Entity"
	x := createEntity()
	if(!reflect.DeepEqual(x.GetName(), createName())) {
		t.Errorf(testedInterface + " doesn't return Name")
	}
	if(!reflect.DeepEqual(x.GetCost(), createCost())) {
		t.Errorf(testedInterface + " doesn't return Cost")
	}
	if(!reflect.DeepEqual(x.GetCastSpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return SpeedModifier")
	}
	if(!reflect.DeepEqual(x.GetTags(), []Tag{createTag()})) {
		t.Errorf(testedInterface + " doesn't return Tags")
	}
	if(!reflect.DeepEqual(x.GetText(), createText())) {
		t.Errorf(testedInterface + " doesn't return Text")
	}
	if(!reflect.DeepEqual(x.GetAbility(), createActivatedAbility())) {
		t.Errorf(testedInterface + " doesn't return Ability")
	}
	if(!reflect.DeepEqual(x.GetAbilitySpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return AbilitySpeed")
	}
	if(!reflect.DeepEqual(x.GetHealth(), createHealth())) {
		t.Errorf(testedInterface + " doesn't return Health")
	}
	if(!reflect.DeepEqual(x.GetAttack(), createAttack())) {
		t.Errorf(testedInterface + " doesn't return Attack")
	}
	success(testedInterface)
}

func TestField(t *testing.T) {
	testedInterface := "Field"
	x := createField()
	if(!reflect.DeepEqual(x.GetName(), createName())) {
		t.Errorf(testedInterface + " doesn't return Name")
	}
	if(!reflect.DeepEqual(x.GetCost(), createCost())) {
		t.Errorf(testedInterface + " doesn't return Cost")
	}
	if(!reflect.DeepEqual(x.GetCastSpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return SpeedModifier")
	}
	if(!reflect.DeepEqual(x.GetTags(), []Tag{createTag()})) {
		t.Errorf(testedInterface + " doesn't return Tags")
	}
	if(!reflect.DeepEqual(x.GetText(), createText())) {
		t.Errorf(testedInterface + " doesn't return Text")
	}
	if(!reflect.DeepEqual(x.GetAbility(), createActivatedAbility())) {
		t.Errorf(testedInterface + " doesn't return Ability")
	}
	if(!reflect.DeepEqual(x.GetAbilitySpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return AbilitySpeed")
	}
	if(!reflect.DeepEqual(x.GetHealth(), createHealth())) {
		t.Errorf(testedInterface + " doesn't return Health")
	}
	success(testedInterface)
}

func TestHeadquarter(t *testing.T) {
	testedInterface := "Headquarter"
	x := createHeadquarter()
	if(!reflect.DeepEqual(x.GetUniqueName(), createUniqueName())) {
		t.Errorf(testedInterface + " doesn't return Name")
	}
	if(!reflect.DeepEqual(x.GetName(), NewName(createUniqueName().GetStringVal()))) {
		t.Errorf(testedInterface + " doesn't return Name")
	}
	if(!reflect.DeepEqual(x.GetTags(), []Tag{createTag()})) {
		t.Errorf(testedInterface + " doesn't return Tags")
	}
	if(!reflect.DeepEqual(x.GetText(), createText())) {
		t.Errorf(testedInterface + " doesn't return Text")
	}
	if(!reflect.DeepEqual(x.GetAbility(), createActivatedAbility())) {
		t.Errorf(testedInterface + " doesn't return Ability")
	}
	if(!reflect.DeepEqual(x.GetAbilitySpeed(), createSpeedModifier())) {
		t.Errorf(testedInterface + " doesn't return AbilitySpeed")
	}
	if(!reflect.DeepEqual(x.GetHealth(), createHealth())) {
		t.Errorf(testedInterface + " doesn't return Health")
	}
	success(testedInterface)
}

func createAction() Action {
	return NewAction(createName(), createCost(), createSpeedModifier(), []Tag{createTag()}, createText(), createEffect())
}

func createEntity() Entity {
	return NewEntity(createName(), createCost(), createSpeedModifier(), []Tag{createTag()}, createText(), createActivatedAbility(), createSpeedModifier(), createHealth(), createAttack())
}

func createField() Field {
	return NewField(createName(), createCost(), createSpeedModifier(), []Tag{createTag()}, createText(), createActivatedAbility(), createSpeedModifier(), createHealth())
}

func createHeadquarter() Headquarter {
	return NewHeadquarter(createUniqueName(), []Tag{createTag()}, createText(), createActivatedAbility(), createSpeedModifier(), createHealth())
}


func success(s string) {
	fmt.Println("[SUCCESS] " + s + " tested")
}
