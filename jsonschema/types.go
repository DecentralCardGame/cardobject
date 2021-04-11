package jsonschema

type validateable interface {
	Validate() error
}

type classBound interface {
	Classes() []string
}

type structType interface {
	validateable
	ValidateStruct() error
	InteractionText() string
}

type interfaceType interface {
	validateable
	ValidateInterface() error
}

type arrayType interface {
	validateable
	ValidateArray() error
	MinMaxItems() (int, int)
	ItemName() string
}

type enumType interface {
	validateable
	ValidateEnum() error
	EnumValues() []string
}

type stringType interface {
	validateable
	ValidateString() error
	MinMaxLength() (int, int)
}

type intType interface {
	validateable
	ValidateInt() error
	MinMax() (int, int)
}

type boolType interface {
	validateable
	ValidateBool() error
}
