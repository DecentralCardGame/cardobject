package cardobject

import (
	"encoding/json"
	"regexp"
	"testing"
)

func TestKeywordedSchemaInteractionFields(t *testing.T) {
	raw, err := KeywordedSchema()
	if err != nil {
		t.Fatal(err)
	}
	var schema struct {
		Definitions map[string]struct {
			InteractionText string                     `json:"interactionText"`
			Children        map[string]json.RawMessage `json:"children"`
		} `json:"definitions"`
	}
	if err := json.Unmarshal([]byte(raw), &schema); err != nil {
		t.Fatal(err)
	}
	placeholders := regexp.MustCompile(`§([A-Za-z]+)`)
	for name, definition := range schema.Definitions {
		for _, match := range placeholders.FindAllStringSubmatch(definition.InteractionText, -1) {
			if _, ok := definition.Children[match[1]]; !ok {
				t.Errorf("%s: interaction refers to absent field %s", name, match[1])
			}
		}
	}
}
