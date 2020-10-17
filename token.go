package cardobject

type token struct {
	Name   string `json:",omitempty"`
	Attack int
	Health int
	Tags   []string `json:",omitempty"`
}

func (t *token) validate() error {
	errorRange := []error{}
	errorRange = append(errorRange, validateCardName(t.Name))
	errorRange = append(errorRange, validateAttack(t.Attack))
	errorRange = append(errorRange, validateHealth(t.Health))
	errorRange = append(errorRange, validateTags(t.Tags))
	return combineErrors(errorRange)
}
