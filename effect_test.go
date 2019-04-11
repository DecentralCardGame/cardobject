package cardobject

import "testing"
import "reflect"

func TestIntManipulation(t *testing.T) {
	testedInterface := "IntManipulation"
	x := createIntManipulation()
	if(!reflect.DeepEqual(x.GetCardSelector(), createBasicSelector())) {
		t.Errorf(testedInterface + " doesn't return CardSelector")
	}
	if(!reflect.DeepEqual(x.GetIntManipulation(), createIntConst())) {
		t.Errorf(testedInterface + " doesn't return ManipulationValue")
	}
	if(x.GetTargetIntPropertyId() != ATTACK) {
		t.Errorf(testedInterface + " doesn't return TargetPropertyId")
	}
	success(testedInterface)
}

func TestStringManipulation(t *testing.T) {
	testedInterface := "StringManipulation"
	x := createStringManipulation()
	if(!reflect.DeepEqual(x.GetCardSelector(), createBasicSelector())) {
		t.Errorf(testedInterface + " doesn't return CardSelector")
	}
	if(!reflect.DeepEqual(x.GetStringManipulation(), createStringConst())) {
		t.Errorf(testedInterface + " doesn't return ManipulationValue")
	}
	if(x.GetTargetStringPropertyId() != TAG) {
		t.Errorf(testedInterface + " doesn't return TargetPropertyId")
	}
	success(testedInterface)
}

func TestEffect(t *testing.T) {
	testedInterface := "Effect"
	x:= createEffect()
	if(!reflect.DeepEqual(x.GetZoneChanges(), []ZoneChange{createZoneChange()})) {
		t.Errorf(testedInterface + " doesn't return ZoneChanges")
	}
	if(!reflect.DeepEqual(x.GetManipulations(), []Manipulation{createIntManipulation()})) {
		t.Errorf(testedInterface + " doesn't return Manipulations")
	}
	if(!reflect.DeepEqual(x.GetProduction(), []Ressource{ENERGY})) {
		t.Errorf(testedInterface + " doesn't return Production")
	}
	success(testedInterface)
}

func createIntManipulation() IntManipulation {
	return NewIntManipulation(createBasicSelector(), createIntConst(), ADD, ATTACK)
}

func createStringManipulation() StringManipulation {
	return NewStringManipulation(createBasicSelector(), createStringConst(), TAG)
}

func createEffect() Effect {
	return NewEffect([]ZoneChange{createZoneChange()}, []Manipulation{createIntManipulation()}, []Ressource{ENERGY})
}