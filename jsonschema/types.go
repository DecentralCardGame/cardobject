package jsonschema

type Validateable interface {
	ValidateType(RootElement) error
}

type RootElement interface {
	Validate() error
	ValidateClasses([]Class) error
}

type classBound interface {
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
	EnumValues() []string
}

type stringType interface {
	Validateable
	MinMaxLength() (int, int)
}

type intType interface {
	Validateable
	MinMax() (int, int)
}

type boolType interface {
	Validateable
	Default() bool
}
