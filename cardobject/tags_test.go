package cardobject

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/DecentralCardGame/cardobject/jsonschema"
)

func TestCanonicalTags(t *testing.T) {
	want := []string{
		"AIR", "ALCHEMIST", "ANIMAL", "ARTIFACT", "ASSASSIN", "BEAST", "BELIEVER", "BOT",
		"BUILDING", "COLOSSUS", "DISASTER", "DOMESTIC", "DRAGON", "DWARF", "EARTH", "ELF",
		"EPIC", "EVENT", "FACTORY", "FIRE", "GOBLIN", "HUMAN", "KNOWLEDGE", "LANDSCAPE",
		"PLANT", "REINFORCEMENT", "WIZARD", "SHRINE", "SPIRIT", "TACTIC", "TECHNOCRAT",
		"UNDEAD", "VEHICLE", "WARRIOR", "WATER", "WEAPON", "WORKER",
	}
	if got := (Tag("")).EnumValues(); !reflect.DeepEqual(got, want) {
		t.Fatalf("tags = %v, want %v", got, want)
	}
	for _, value := range want {
		tag := Tag(value)
		if err := tag.ValidateType(nil); err != nil {
			t.Errorf("tag %s: %v", value, err)
		}
	}
	for _, value := range []Tag{"DWARVEN", "Dwarven", "dwarf", "", "INVALID"} {
		if err := value.ValidateType(nil); err == nil {
			t.Errorf("accepted noncanonical tag %q", value)
		}
	}
	tags := Tags{"DWARF", "GOBLIN", "ELF"}
	if err := tags.ValidateType(nil); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(tags)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `["DWARF","GOBLIN","ELF"]` {
		t.Fatalf("unexpected tag JSON: %s", data)
	}
	var decoded Tags
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, tags) {
		t.Fatalf("tags did not round-trip: %v", decoded)
	}
	schema, err := jsonschema.Reflect(&tags)
	if err != nil {
		t.Fatal(err)
	}
	if got := schema.Definitions["Tag"].Enum; !reflect.DeepEqual(got, want) {
		t.Fatalf("schema tags = %v, want %v", got, want)
	}
}
