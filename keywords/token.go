package keywords

import (
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

type token struct {
	Avatar     *avatar     `json:",omitempty"`
	Beast      *beast      `json:",omitempty"`
	Bot        *bot        `json:",omitempty"`
	Powerstone *powerstone `json:",omitempty"`
	Recruit    *recruit    `json:",omitempty"`
	Spirit     *spirit     `json:",omitempty"`
}

func (t token) ValidateType(r jsonschema.RootElement) error {
	implementer, err := t.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

func (t token) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(t)
}
