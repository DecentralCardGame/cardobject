package keywords

import (
	"errors"
	"strconv"

	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/jsonschema"
)

const maxKeywords int = 8

type CardRootValidator struct {
	Keywords      int
	ClassProvider cardobject.ClassProvider
	Type          string
}

// Validate Ensures that the card was built correctly and returns errors otherwise
func (c CardRootValidator) Validate() error {
	return nil
}

// ValidateClasses Checks if the given classes are covered by the card
func (c CardRootValidator) ValidateClasses(cb jsonschema.ClassBound) error {
	if c.ClassProvider.ClassRestriction().Contains(cb.Classes()) {
		return nil
	}
	return errors.New("Classes are not covered by the cards classes")
}

func (c *CardRootValidator) ValidateKeywordCount() error {
	c.Keywords = c.Keywords + 1
	if c.Keywords <= maxKeywords {
		return nil
	}
	return errors.New("too many keywords, max is " + strconv.Itoa(maxKeywords))
}

// ValidateTarget Checks if the given classes are covered by the card
func (c CardRootValidator) ValidateTarget(t jsonschema.Targeting) error {
	ty, tm := t.Targets()
	if tm == "THIS" {
		for _, v := range ty {
			if v == c.Type {
				return nil
			}
		}
		return errors.New("Class " + c.Type + " is not covered by the target classes")
	} else {
		return nil
	}
}

// Card Represents the data-structure of a crowd control card in the blockchain
type Card struct {
	Action      *action      `json:",omitempty"`
	Entity      *entity      `json:",omitempty"`
	Place       *place       `json:",omitempty"`
	Headquarter *headquarter `json:",omitempty"`
}

// Validate Ensures that the card was built correctly and returns errors otherwise
func (c Card) Validate() error {
	classProvider, _ := c.FindImplementer()
	cardRoot := CardRootValidator{0, classProvider.(cardobject.ClassProvider), c.GetType()}
	return c.ValidateType(&cardRoot)
}

// ValidateType Ensures that the type "Card" is build correctly in the context of the RootElement
func (c Card) ValidateType(r jsonschema.RootElement) error {
	implementer, err := c.FindImplementer()
	if err != nil {
		return err
	}
	return implementer.ValidateType(r)
}

// FindImplementer Returns which of its children "implement" the type "Card"
func (c Card) FindImplementer() (jsonschema.Validateable, error) {
	return jsonschema.FindImplementer(c)
}

func (c Card) GetType() string {
	if c.Action != nil {
		return cardobject.ACTIONTYPE
	}
	if c.Entity != nil {
		return cardobject.ENTITYTYPE
	}
	if c.Place != nil {
		return cardobject.PLACETYPE
	}
	return cardobject.HEADQUARTERTYPE
}

type action struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Effects        effects
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	Rarity         *cardobject.Rarity `json:",omitempty"`
}

func (a action) ClassRestriction() cardobject.Class {
	return a.Class
}

func (a action) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(a, r)
}

func (a action) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Effects §FlavourText §Tags §Keywords §RulesTexts §Rarity"
}

type entity struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Abilities      abilities
	Attack         cardobject.Attack
	Health         cardobject.Health
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	Rarity         *cardobject.Rarity `json:",omitempty"`
}

func (e entity) ClassRestriction() cardobject.Class {
	return e.Class
}

func (e entity) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(e, r)
}

func (e entity) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Attack §Health §FlavourText §Tags §Keywords §RulesTexts §Rarity"
}

type place struct {
	CardName       cardobject.CardName
	CastingCost    cardobject.CastingCost
	AdditionalCost *cardobject.AdditionalCost `json:",omitempty"`
	Class          cardobject.Class
	Abilities      abilities
	Health         cardobject.Health
	FlavourText    cardobject.FlavourText
	Tags           cardobject.Tags
	Keywords       cardobject.Keywords
	Rarity         *cardobject.Rarity `json:",omitempty"`
}

func (p place) ClassRestriction() cardobject.Class {
	return p.Class
}

func (p place) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(p, r)
}

func (p place) InteractionText() string {
	return "§CardName §CastingCost §AdditionalCost §Class §Abilities §Health §FlavourText §Tags §Keywords §RulesTexts §Rarity"
}

type headquarter struct {
	CardName    cardobject.CardName
	Class       cardobject.Class
	Delay       cardobject.Delay
	Abilities   abilities
	Health      cardobject.Health
	FlavourText cardobject.FlavourText
	Tags        cardobject.Tags
	Keywords    cardobject.Keywords
	Rarity      *cardobject.Rarity `json:",omitempty"`
}

func (h headquarter) ClassRestriction() cardobject.Class {
	return h.Class
}

func (h headquarter) ValidateType(r jsonschema.RootElement) error {
	return jsonschema.ValidateStruct(h, r)
}

func (h headquarter) InteractionText() string {
	return "§CardName §Class §Delay §Abilities §Health §Growth §StartingHandSize §Wisdom §FlavourText §Tags §Keywords §RulesTexts §Rarity"
}
