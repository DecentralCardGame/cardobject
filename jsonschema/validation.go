package jsonschema

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var classBoundElem = reflect.TypeOf((*ClassBound)(nil)).Elem()

func validate(i interface{}, r RootElement) error {
	t := reflect.TypeOf(i)

	if t.Implements(structTypeElem) {
		return ValidateStruct(i.(structType), r)
	}
	if t.Implements(interfaceTypeElem) {
		return validateInterface(i.(interfaceType), r)
	}
	if t.Implements(arrayTypeElem) {
		return validateArray(i.(arrayType), r)
	}
	if t.Implements(enumTypeElem) {
		return validateEnum(i.(enumType), r)
	}
	if t.Implements(stringTypeElem) {
		return validateString(i.(stringType), r)
	}
	if t.Implements(intTypeElem) {
		return validateInt(i.(intType), r)
	}
	if t.Implements(boolTypeElem) {
		return validateBool(i.(boolType), r)
	}

	return errors.New("Unknown Type " + t.Name())
}

func ValidateStruct(s structType, r RootElement) error {
	errorRange := []error{}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if t.Implements(classBoundElem) {
		ClassBound := s.(ClassBound)
		errorRange = append(errorRange, r.ValidateClasses(ClassBound))
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() != reflect.Ptr || !field.IsNil() {
			fieldValue := field.Interface().(Validateable)
			errorRange = append(errorRange, validate(fieldValue, r)) //fieldValue.ValidateType(r))
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

func validateInterface(i interfaceType, r RootElement) error {
	implementer, err := i.FindImplementer()
	if err != nil {
		return err
	}
	return validate(implementer, r)
}

func validateArray(a arrayType, r RootElement) error {
	array := reflect.ValueOf(a)
	arrayLenght := array.Len()
	errorRange := []error{}
	for i := 0; i < arrayLenght; i++ {
		err := validate(array.Index(i), r)
		if err != nil {
			errorRange = append(errorRange, err)
		}
	}
	return CombineErrors(errorRange)
}

func validateEnum(e enumType, r RootElement) error {
	values := e.EnumValues()
	name := reflect.TypeOf(e).Name()
	s := e.String()
	for _, v := range values {
		if v == s {
			return nil
		}
	}
	return errors.New(name + " must be one of: " + strings.Join(values, ","))
}

func validateString(s stringType, r RootElement) error {
	minLength, maxLength := s.MinMaxLength()
	name := reflect.TypeOf(s).Name()
	length := len(s.String())
	if length < minLength || length > maxLength {
		return errors.New(name + " must be between " + strconv.Itoa(minLength) + " and " + strconv.Itoa(maxLength) + " characters long")
	}
	return nil
}

func validateInt(i intType, r RootElement) error {
	min, max := i.MinMax()
	name := reflect.TypeOf(i).Name()
	if i.Int() < min || i.Int() > max {
		return errors.New(name + " must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}

func validateBool(b boolType, r RootElement) error {
	return nil
}

//FindImplementer returns the non-nil field of a struct if it has exactly one
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

//BasicEnum EnumTypes should be derived from BasicEnum
type BasicEnum string

//BasicString StringTypes should be derived from BasicString
type BasicString string

//BasicInt IntTypes should be derived from BasicInt
type BasicInt int

//BasicBool BoolTypes should be derived from BasicBool
type BasicBool bool

//Class ClassTypes should be derived from Class
type Class string

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
