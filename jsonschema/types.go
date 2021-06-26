package jsonschema

//Validateable Every Type of the struct tree must be Validateable
type Validateable interface {
	ValidateType(RootElement) error
}

//RootElement Root of the struct tree and entrypoint of the schema
type RootElement interface {
	Validate() error
	ValidateClasses(ClassBound) error
}

//ClassBound A Type that has a class restriction
type ClassBound interface {
	Classes() []Class
}

type structType interface {
	Validateable
	InteractionText() string
}

type interfaceType interface {
	Validateable
	FindImplementer() (Validateable, error)
}

type arrayType interface {
	Validateable
	MinMaxItems() (int, int)
	ItemName() string
}

type enumType interface {
	Validateable
	String() string
	EnumValues() []string
}

type stringType interface {
	Validateable
	String() string
	MinMaxLength() (int, int)
}

type intType interface {
	Validateable
	Int() int
	MinMax() (int, int)
}

type boolType interface {
	Validateable
	Bool() bool
	Default() bool
}
