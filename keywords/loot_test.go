package keywords

import (
	"encoding/json"
	"testing"
)

func TestLootDescription(t *testing.T) {
	expected := "If at least one of your attackers dealt damage to an opposing Place or HQ, activate Effects."
	if actual := (loot{}).Description(); actual != expected {
		t.Errorf("Loot description = %q, want %q", actual, expected)
	}
}

func TestLootSchemaDescription(t *testing.T) {
	data, err := Schema()
	if err != nil {
		t.Fatal(err)
	}
	var schema struct {
		Definitions map[string]struct {
			Description string `json:"description"`
		} `json:"definitions"`
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatal(err)
	}
	expected := "If at least one of your attackers dealt damage to an opposing Place or HQ, activate Effects."
	if actual := schema.Definitions["loot"].Description; actual != expected {
		t.Errorf("Loot schema description = %q, want %q", actual, expected)
	}
}

func TestLoot(t *testing.T) {
	loot := loot{nil}
	err := loot.ValidateType(allClassesTestCard())
	if err != nil {
		t.Error(err)
	}
}
