package jsonschema

import (
	"encoding/json"

	"github.com/iancoleman/orderedmap"
)

func (s *Schema) MarshalJSONForJS(name string) ([]byte, error) {
	wrappedType := orderedmap.New()
	wrappedType.Set(name, s.Type)
	b, err := json.Marshal(wrappedType)
	if err != nil {
		return nil, err
	}
	if s.Definitions == nil || len(s.Definitions) == 0 {
		return b, nil
	}
	d, err := json.Marshal(struct {
		Definitions Definitions `json:"definitions,omitempty"`
	}{s.Definitions})
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return d, nil
	} else {
		b[len(b)-1] = ','
		return append(b, d[1:]...), nil
	}
}

func (t *Type) MarshalJSON() ([]byte, error) {
	type Type_ Type
	b, err := json.Marshal((*Type_)(t))
	if err != nil {
		return nil, err
	}
	if t.Extras == nil || len(t.Extras) == 0 {
		return b, nil
	}
	m, err := json.Marshal(t.Extras)
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return m, nil
	} else {
		b[len(b)-1] = ','
		return append(b, m[1:]...), nil
	}
}
