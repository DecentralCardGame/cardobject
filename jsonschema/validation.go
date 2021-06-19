package jsonschema

import (
	"errors"
	"reflect"
	"strings"
)

func ValidateStruct(s structType, r RootElement) error {
	errorRange := []error{}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() != reflect.Ptr || !field.IsNil() {
			fieldValue := field.Interface().(Validateable)
			errorRange = append(errorRange, fieldValue.Validate(r))
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
	if implementer == nil || error != nil {
		return nil, errors.New(typeOfB.Name() + " implemented by not exactly one option")
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

type BasicEnum string

type BasicString string

type BasicInt int

type BasicBool bool

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
