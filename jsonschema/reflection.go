package jsonschema

import (
	"errors"
	"reflect"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/iancoleman/orderedmap"
)

// Schema is the root schema.
type Schema struct {
	*Type
	Definitions Definitions
}

// Type represents a JSON Schema object type.
type Type struct {
	//Contained in all types
	Ref         string                 `json:"$ref,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Definitions Definitions            `json:"definitions,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Properties  *orderedmap.OrderedMap `json:"children,omitempty"`
	Classes     []Class                `json:"classes,omitempty"`
	Extras      map[string]interface{} `json:"-"`

	//exclusive for structs
	Required        []string `json:"required,omitempty"`
	InteractionText string   `json:"interactionText,omitempty"`

	//exclusive for enums
	Enum []string `json:"enum,omitempty"`

	//exclusive for strings
	Pattern string `json:"pattern,omitempty"`

	//exclusive for ints, strings, arrays
	Max int64 `json:"max,omitempty"`
	Min int64 `json:"min,omitempty"`
}

// Reflect reflects to Schema from a value using the default Reflector
func Reflect(v interface{}) (*Schema, error) {
	return ReflectFromType(reflect.TypeOf(v))
}

// ReflectFromType generates root schema using the default Reflector
func ReflectFromType(t reflect.Type) (*Schema, error) {
	r := &Reflector{}
	return r.ReflectFromType(t)
}

// A Reflector reflects values into a Schema.
type Reflector struct {
	// Use package paths as well as type names, to avoid conflicts.
	// Without this setting, if two packages contain a type with the same name,
	// and both are present in a schema, they will conflict and overwrite in
	// the definition map and produce bad output.  This is particularly
	// noticeable when using DoNotReference.
	FullyQualifyTypeNames bool
}

// Reflect reflects to Schema from a value.
func (r *Reflector) Reflect(v interface{}) (*Schema, error) {
	return r.ReflectFromType(reflect.TypeOf(v))
}

// ReflectFromType generates root schema
func (r *Reflector) ReflectFromType(t reflect.Type) (*Schema, error) {
	definitions := Definitions{}
	st, err := r.reflectTypeToSchema(definitions, t)
	if err != nil {
		return nil, err
	}
	s := &Schema{
		Type:        st,
		Definitions: definitions,
	}
	return s, nil
}

// Definitions hold schema definitions
type Definitions map[string]*Type

//needed for implements check
var structTypeElem = reflect.TypeOf((*structType)(nil)).Elem()
var interfaceTypeElem = reflect.TypeOf((*interfaceType)(nil)).Elem()
var arrayTypeElem = reflect.TypeOf((*arrayType)(nil)).Elem()
var enumTypeElem = reflect.TypeOf((*enumType)(nil)).Elem()
var stringTypeElem = reflect.TypeOf((*stringType)(nil)).Elem()
var intTypeElem = reflect.TypeOf((*intType)(nil)).Elem()
var boolTypeElem = reflect.TypeOf((*boolType)(nil)).Elem()

//dummy for method-calls
var dummyInputs = make([]reflect.Value, 0)

func (r *Reflector) reflectTypeToSchema(definitions Definitions, t reflect.Type) (*Type, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// Already added to definitions?
	if _, ok := definitions[r.typeName(t)]; ok {
		return &Type{Ref: "#/definitions/" + r.typeName(t)}, nil
	}

	if t.Implements(structTypeElem) {
		str, err := r.reflectStruct(definitions, t)
		return str, err
	}
	if t.Implements(interfaceTypeElem) {
		return r.reflectInterface(definitions, t)
	}
	if t.Implements(arrayTypeElem) {
		return r.reflectArray(definitions, t)
	}
	if t.Implements(enumTypeElem) {
		return r.reflectEnum(definitions, t), nil
	}
	if t.Implements(stringTypeElem) {
		return r.reflectString(definitions, t), nil
	}
	if t.Implements(intTypeElem) {
		return r.reflectInt(definitions, t), nil
	}
	if t.Implements(boolTypeElem) {
		return r.reflectBool(definitions, t), nil
	}
	return nil, errors.New("Unknown Type " + t.Name())
}

func (r *Reflector) typeName(t reflect.Type) string {
	if r.FullyQualifyTypeNames {
		return t.PkgPath() + "." + t.Name()
	}
	return t.Name()
}

// Refects a struct to a JSON Schema type.
func (r *Reflector) reflectStruct(definitions Definitions, t reflect.Type) (*Type, error) {
	defType := "struct"
	name := getReadableFromCamelCase(t.Name())

	interactionText := reflect.New(t).MethodByName("InteractionText").Call(dummyInputs)[0].String()

	descriptionMethod := reflect.New(t).MethodByName("Description")

	description := "Build a " + name
	if descriptionMethod.IsValid() {
		description = descriptionMethod.Call(dummyInputs)[0].String()
	}

	st := &Type{
		Name:            name,
		Type:            defType,
		Properties:      orderedmap.New(),
		Description:     description,
		InteractionText: interactionText,
	}
	definitions[r.typeName(t)] = st
	err := r.reflectStructFields(st, definitions, t)

	var classBoundElem = reflect.TypeOf((*classBound)(nil)).Elem()

	if t.Implements(classBoundElem) {
		classValues := reflect.New(t).MethodByName("Classes").Call(dummyInputs)[0]
		var classes, _ = classValues.Interface().([]Class)

		st.Classes = classes
	}

	if err != nil {
		return nil, err
	}

	textErr := r.validateInteractionText(st)

	if textErr != nil {
		return nil, textErr
	}
	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}, nil
}

func (r *Reflector) reflectStructFields(st *Type, definitions Definitions, t reflect.Type) error {
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, exist, required := r.reflectFieldName(f)
		// if anonymous and exported type should be processed recursively
		// current type should inherit properties of anonymous one
		if name == "" {
			if f.Anonymous && !exist {
				r.reflectStructFields(st, definitions, f.Type)
			}
			continue
		}

		property, err := r.reflectTypeToSchema(definitions, f.Type)
		if err != nil {
			return err
		}
		st.Properties.Set(name, property)

		if required {
			st.Required = append(st.Required, name)
		}
	}
	return nil
}

func (r *Reflector) validateInteractionText(st *Type) error {
	text := st.InteractionText
	children := st.Properties.Keys()
	for _, name := range children {
		if !strings.Contains(text, "§"+name) {
			return errors.New(st.Name + "'s interactionText doesn't contain §" + name)
		}
	}
	return nil
}

// Refects a struct to a JSON Schema type.
func (r *Reflector) reflectInterface(definitions Definitions, t reflect.Type) (*Type, error) {
	defType := "interface"
	name := getReadableFromCamelCase(t.Name())

	st := &Type{
		Name:        name,
		Type:        defType,
		Properties:  orderedmap.New(),
		Description: "Choose a " + name,
	}
	definitions[r.typeName(t)] = st
	err := r.reflectInterfaceFields(st, definitions, t)

	if err != nil {
		return nil, err
	}

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}, nil
}

func (r *Reflector) reflectInterfaceFields(st *Type, definitions Definitions, t reflect.Type) error {
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, exist, _ := r.reflectFieldName(f)
		// if anonymous and exported type should be processed recursively
		// current type should inherit properties of anonymous one
		if name == "" {
			if f.Anonymous && !exist {
				r.reflectStructFields(st, definitions, f.Type)
			}
			continue
		}

		property, err := r.reflectTypeToSchema(definitions, f.Type)
		if err != nil {
			return err
		}
		st.Properties.Set(name, property)
	}
	return nil
}

func (r *Reflector) reflectFieldName(f reflect.StructField) (string, bool, bool) {
	jsonTags, exist := f.Tag.Lookup("json")
	if !exist {
		jsonTags = f.Tag.Get("yaml")
	}

	jsonTagsList := strings.Split(jsonTags, ",")

	if ignoredByJSONTags(jsonTagsList) {
		return "", exist, false
	}

	jsonSchemaTags := strings.Split(f.Tag.Get("jsonschema"), ",")
	if ignoredByJSONSchemaTags(jsonSchemaTags) {
		return "", exist, false
	}

	name := f.Name
	required := requiredFromJSONTags(jsonTagsList)

	if jsonTagsList[0] != "" {
		name = jsonTagsList[0]
	}

	// field not anonymous and not export has no export name
	if !f.Anonymous && f.PkgPath != "" {
		name = ""
	}

	// field anonymous but without json tag should be inherited by current type
	if f.Anonymous && !exist {
		name = ""
	}

	return name, exist, required
}

func requiredFromJSONTags(tags []string) bool {
	for _, tag := range tags[1:] {
		if tag == "omitempty" {
			return false
		}
	}
	return true
}

func ignoredByJSONTags(tags []string) bool {
	return tags[0] == "-"
}

func ignoredByJSONSchemaTags(tags []string) bool {
	return tags[0] == "-"
}

func (r *Reflector) reflectArray(definitions Definitions, t reflect.Type) (*Type, error) {
	defType := "array"
	name := getReadableFromCamelCase(t.Name())

	returnType := &Type{
		Name:        name,
		Type:        defType,
		Properties:  orderedmap.New(),
		Description: "Add some " + name,
	}

	//type has MinMaxItems function because it implements arrayType
	minMaxItems := reflect.New(t).MethodByName("MinMaxItems").Call(dummyInputs)
	returnType.Min = minMaxItems[0].Int()
	returnType.Max = minMaxItems[1].Int()

	property, err := r.reflectTypeToSchema(definitions, t.Elem())

	if err != nil {
		return nil, err
	}

	//remove plural-'s' for singular children
	itemName := reflect.New(t).MethodByName("ItemName").Call(dummyInputs)[0].String()

	returnType.Properties.Set(itemName, property)

	definitions[r.typeName(t)] = returnType

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}, nil
}

func (r *Reflector) reflectEnum(definitions Definitions, t reflect.Type) *Type {
	defType := "enum"

	returnType := &Type{
		Name: getReadableFromCamelCase(t.Name()),
		Type: defType,
	}

	//type has EnumValues function because it implements enumType
	enumValues := reflect.New(t).MethodByName("EnumValues").Call(dummyInputs)[0]
	returnType.Enum, _ = enumValues.Interface().([]string)

	definitions[r.typeName(t)] = returnType

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}
}

func (r *Reflector) reflectString(definitions Definitions, t reflect.Type) *Type {
	defType := "string"

	returnType := &Type{
		Name: getReadableFromCamelCase(t.Name()),
		Type: defType,
	}

	//type has MinMaxLength function because it implements stringType
	minMaxLength := reflect.New(t).MethodByName("MinMaxLength").Call(dummyInputs)
	returnType.Min = minMaxLength[0].Int()
	returnType.Max = minMaxLength[1].Int()

	definitions[r.typeName(t)] = returnType

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}
}

func (r *Reflector) reflectInt(definitions Definitions, t reflect.Type) *Type {
	defType := "int"

	returnType := &Type{
		Name: getReadableFromCamelCase(t.Name()),
		Type: defType,
	}

	//type has MinMax function because it implements intType
	minMax := reflect.New(t).MethodByName("MinMax").Call(dummyInputs)
	returnType.Min = minMax[0].Int()
	returnType.Max = minMax[1].Int()

	definitions[r.typeName(t)] = returnType

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}
}

func (r *Reflector) reflectBool(definitions Definitions, t reflect.Type) *Type {
	defType := "bool"

	returnType := &Type{
		Name: getReadableFromCamelCase(t.Name()),
		Type: defType,
	}

	definitions[r.typeName(t)] = returnType

	return &Type{
		Ref: "#/definitions/" + r.typeName(t),
	}
}

func getReadableFromCamelCase(s string) string {
	readable := ""
	words := camelcase.Split(s)
	for i, word := range words {
		if i > 0 {
			readable += " "
		}
		readable += strings.Title(word)
	}
	return readable
}
