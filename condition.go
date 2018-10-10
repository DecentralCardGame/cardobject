package cardobject

import "reflect"
import "errors"

type Condition interface {
	GetComparator() Comparator
}

type CardCondition interface {
	GetComparator() Comparator
	IsCardCondition() bool
}

type PlayerCondition interface {
	GetComparator() Comparator
	IsPlayerCondition() bool
}

type IntCondition interface {
	GetComparator() Comparator
	ExtractIntCond() int
	ExtractIntProp() int
}

type StringCondition interface {
	GetComparator() Comparator
	ExtractStringCond() string
	ExtractStringProp() string
}

func NewIntCondition(prop IntProperty, comp Comparator, val int) (IntCondition, error) {
	cpType := reflect.TypeOf((*CardIntProperty)(nil)).Elem()
	ppType := reflect.TypeOf((*PlayerIntProperty)(nil)).Elem()
	t := reflect.TypeOf(prop)

	switch {
		case t.Implements(cpType): return &cardIntCondition{&cardCondition{&condition{comp}}, prop, val}, nil
		case t.Implements(ppType): return &playerIntCondition{&playerCondition{&condition{comp}}, prop, val}, nil
		default: return &cardIntCondition{}, errors.New("unknown Type")
	}
}

func NewStringCondition(prop StringProperty, val string) (StringCondition, error) {
	cpType := reflect.TypeOf((*CardStringProperty)(nil)).Elem()
	ppType := reflect.TypeOf((*PlayerStringProperty)(nil)).Elem()
	t := reflect.TypeOf(prop)

	switch {
		case t.Implements(cpType): return &cardStringCondition{&cardCondition{&condition{EQUAL}}, prop, val}, nil
		case t.Implements(ppType): return &playerStringCondition{&playerCondition{&condition{EQUAL}}, prop, val}, nil
		default: return &cardStringCondition{}, errors.New("unknown Type")
	}
}

type condition struct {
	comparator Comparator
}

type cardCondition struct {
	*condition
}
type playerCondition struct {
	*condition
}

type cardIntCondition struct {
	*cardCondition
	prop IntProperty
	value int
}
type cardStringCondition struct {
	*cardCondition
	prop StringProperty
	value string
}
type playerIntCondition struct {
	*playerCondition
	prop IntProperty
	value int
}
type playerStringCondition struct {
	*playerCondition
	prop StringProperty
	value string
}

type noneCondition int

const NOCODITION noneCondition = 1

func (c *condition) GetComparator() Comparator {
	return c.comparator
}

func (cc *cardCondition) IsCardCondition() bool {
	return true
}

func (pc *playerCondition) IsPlayerCondition() bool {
	return true
}

func (nc *noneCondition) IsCardCondition() bool {
	return true
}

func (nc *noneCondition) IsPlayerCondition() bool {
	return true
}

func (cic *cardIntCondition) ExtractIntCond() int {
	return cic.value
}

func (pic *playerIntCondition) ExtractIntCond() int {
	return pic.value
}

func (csc *cardStringCondition) ExtractStringCond() string {
	return csc.value
}

func (psc *playerStringCondition) ExtractStringCond() string {
	return psc.value
}

func (cic *cardIntCondition) ExtractIntProp() int {
	return cic.prop.ExtractIntProp()
}

func (pic *playerIntCondition) ExtractIntProp() int {
	return pic.prop.ExtractIntProp()
}

func (csc *cardStringCondition) ExtractStringProp() string {
	return csc.prop.ExtractStringProp()
}

func (psc *playerStringCondition) ExtractStringProp() string {
	return psc.prop.ExtractStringProp()
}
