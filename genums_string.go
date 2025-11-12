package genums

// EnumString is a specialized enum interface for string-based enumerations.
// It extends BaseEnum with string as the underlying value type, providing
// type-safe string enum functionality.
//
// Type parameters:
//   - T: The concrete enum type that implements this interface
type EnumString[T any] interface {
	BaseEnum[T, string]
}

// enumStringImpl is the internal implementation of the EnumString interface.
// It embeds enumImpl with string as the value type to provide all
// string-specific enum functionality.
//
// Type parameters:
//   - T: The concrete enum type (used for type safety)
type enumStringImpl[T any] struct {
	enumImpl[T, string]
}

// NewEnumString creates a new string-based enum instance with the given value.
// This is the primary constructor for creating string enum values.
//
// Type parameters:
//   - T: The target enum type to create
//
// Parameters:
//   - value: The string value to wrap in the enum
//
// Returns:
//   - A new EnumString instance containing the specified value
//
// Example:
//
//	type Status EnumString[Status]
//	active := NewEnumString[Status]("active")
func NewEnumString[T any](value string) EnumString[T] {
	return enumStringImpl[T]{enumImpl[T, string]{value: value}}
}

// StringFactory is a factory type that provides a convenient way to create
// string-based enum instances. It implements a factory pattern for enums,
// allowing for more flexible instantiation.
//
// Type parameters:
//   - T: The enum type to be created (must implement EnumString[T])
type StringFactory[T EnumString[T]] struct{}

// New creates a new enum instance of type T with the specified string value.
// This method allows the factory to produce properly typed enum instances
// using the underlying NewEnumString constructor.
//
// Parameters:
//   - value: The string value for the new enum instance
//
// Returns:
//   - A new enum instance of type T with the given value
//
// Example:
//
//	factory := NewStringFactory[Status]()
//	active := factory.New("active")
func (f StringFactory[T]) New(value string) T {
	return any(NewEnumString[T](value)).(T)
}

// NewStringFactory creates a new StringFactory instance for the specified enum type.
// This function initializes a factory that can be used to create multiple enum
// instances of the same type.
//
// Type parameters:
//   - T: The enum type the factory will create (must implement EnumString[T])
//
// Returns:
//   - A new StringFactory instance for type T
//
// Example:
//
//	type Status EnumString[Status]
//	factory := NewStringFactory[Status]()
//	active := factory.New("active")
//	inactive := factory.New("inactive")
func NewStringFactory[T EnumString[T]]() StringFactory[T] {
	return StringFactory[T]{}
}
