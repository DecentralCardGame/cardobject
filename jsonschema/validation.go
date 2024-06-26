package jsonschema

import (
	"errors"
	"reflect"
	"strings"
)

var classBoundElem = reflect.TypeOf((*ClassBound)(nil)).Elem()
var targetingElem = reflect.TypeOf((*Targeting)(nil)).Elem()

func ValidateStruct(s structType, r RootElement) error {
	errorRange := []error{}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if t.Implements(classBoundElem) {
		ClassBound := s.(ClassBound)
		errorRange = append(errorRange, r.ValidateClasses(ClassBound))
	}

	if t.Implements(targetingElem) {
		Targeting := s.(Targeting)
		errorRange = append(errorRange, r.ValidateTarget(Targeting))
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() != reflect.Ptr || !field.IsNil() {
			fieldValue := field.Interface().(Validateable)
			errorRange = append(errorRange, fieldValue.ValidateType(r))
		} else {
			fieldType := t.Field(i)
			tags, exists := fieldType.Tag.Lookup("json")
			if !exists {
				errorRange = append(errorRange, errors.New(t.Name()+" requires field "+fieldType.Name))
			} else if requiredFromJSONTags(strings.Split(tags, ",")) {
				errorRange = append(errorRange, errors.New(t.Name()+" requires field "+fieldType.Name))
			}
		}
	}
	return CombineErrors(errorRange)
}

// FindImplementer returns the non-nil field of a struct if it has exactly one
func FindImplementer(i interfaceType) (Validateable, error) {
	valueOfB := reflect.ValueOf(i)
	typeOfB := reflect.TypeOf(i)
	possibleImplementer := []Validateable{}
	for k := 0; k < valueOfB.NumField(); k++ {
		if !typeOfB.Field(k).Anonymous {
			possibleImplementer = append(possibleImplementer, valueOfB.Field(k).Interface().(Validateable))
		}
	}
	implementer, error := getDefinedValidateable(possibleImplementer)
	if implementer == nil {
		return nil, errors.New(typeOfB.Name() + " not implemented")
	}
	if error != nil {
		return nil, error
	}
	return implementer, nil
}

func getDefinedValidateable(possibleImplementer []Validateable) (Validateable, error) {
	var implementer Validateable
	implementerFound := false
	for _, b := range possibleImplementer {
		if !reflect.ValueOf(b).IsNil() {
			if implementerFound {
				return nil, errors.New("ImplementerError")
			}
			implementer = b
			implementerFound = true
		}
	}
	return implementer, nil
}

/*ItemNameFromArray computes the supposed item name of an array
 *by removing the plural 's' of the array name
 */
func ItemNameFromArray(array arrayType) string {
	name := reflect.TypeOf(array).Name()
	itemName := strings.Title(strings.TrimSuffix(name, "s"))
	return itemName
}

// BasicEnum EnumTypes should be derived from BasicEnum
type BasicEnum string

// BasicString StringTypes should be derived from BasicString
type BasicString string

// BasicInt IntTypes should be derived from BasicInt
type BasicInt int

// BasicBool BoolTypes should be derived from BasicBool
type BasicBool bool

// Class ClassTypes should be derived from Class
type Class string

// TargetMode targeting modes should be derived from this class
type TargetMode BasicEnum

func CombineErrors(errorRange []error) error {
	isError := false
	errorString := ""
	for _, error := range errorRange {
		if error != nil {
			isError = true
			errorString = errorString + "\n" + error.Error()
		}
	}
	if isError {
		return errors.New(errorString)
	}
	return nil
}
