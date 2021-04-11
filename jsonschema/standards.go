package jsonschema

import (
	"errors"
	"reflect"
	"strings"
)

func ValidateStruct(s structType) error {
	errorRange := []error{}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() != reflect.Ptr || !field.IsNil() {
			fieldValue := field.Interface().(validateable)
			errorRange = append(errorRange, fieldValue.Validate())
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

func ValidateInterface(i interfaceType) error {
	valueOfB := reflect.ValueOf(i)
	typeOfB := reflect.TypeOf(i)
	possibleImplementer := []validateable{}
	for k := 0; k < valueOfB.NumField(); k++ {
		if !typeOfB.Field(k).Anonymous {
			possibleImplementer = append(possibleImplementer, valueOfB.Field(k).Interface().(validateable))
		}
	}
	implementer, error := xorInterface(possibleImplementer)
	if implementer == nil || error != nil {
		return errors.New(typeOfB.Name() + " implemented by not exactly one option")
	}
	return implementer.Validate()
}

func xorInterface(possibleImplementer []validateable) (validateable, error) {
	var implementer validateable
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
