package genums

// EnumBool is a specialized enum interface for boolean-based enumerations.
// It extends BaseEnum with bool as the underlying value type, providing
// type-safe boolean enum functionality.
//
// Type parameters:
//   - T: The concrete enum type that implements this interface
type EnumBool[T any] interface {
	BaseEnum[T, bool]
}

// enumBoolImpl is the internal implementation of the EnumBool interface.
// It embeds enumImpl with bool as the value type to provide all
// boolean-specific enum functionality.
//
// Type parameters:
//   - T: The concrete enum type (used for type safety)
type enumBoolImpl[T any] struct {
	enumImpl[T, bool]
}

// NewEnumBool creates a new boolean-based enum instance with the given value.
// This is the primary constructor for creating boolean enum values.
//
// Type parameters:
//   - T: The target enum type to create
//
// Parameters:
//   - value: The boolean value to wrap in the enum
//
// Returns:
//   - A new EnumBool instance containing the specified value
//
// Example:
//
//	type FeatureFlag EnumBool[FeatureFlag]
//	enabled := NewEnumBool[FeatureFlag](true)
//	disabled := NewEnumBool[FeatureFlag](false)
func NewEnumBool[T any](value bool) EnumBool[T] {
	return enumBoolImpl[T]{enumImpl[T, bool]{value: value}}
}

// BoolFactory is a factory type that provides a convenient way to create
// boolean-based enum instances. It implements a factory pattern for enums,
// allowing for more flexible instantiation.
//
// Type parameters:
//   - T: The enum type to be created (must implement EnumBool[T])
type BoolFactory[T EnumBool[T]] struct{}

// New creates a new enum instance of type T with the specified boolean value.
// This method allows the factory to produce properly typed enum instances
// using the underlying NewEnumBool constructor.
//
// Parameters:
//   - value: The boolean value for the new enum instance
//
// Returns:
//   - A new enum instance of type T with the given value
//
// Example:
//
//	factory := NewBoolFactory[FeatureFlag]()
//	enabled := factory.New(true)
//	disabled := factory.New(false)
func (f BoolFactory[T]) New(value bool) T {
	return any(NewEnumBool[T](value)).(T)
}

// NewBoolFactory creates a new BoolFactory instance for the specified enum type.
// This function initializes a factory that can be used to create multiple enum
// instances of the same type.
//
// Type parameters:
//   - T: The enum type the factory will create (must implement EnumBool[T])
//
// Returns:
//   - A new BoolFactory instance for type T
//
// Example:
//
//	type FeatureFlag EnumBool[FeatureFlag]
//	factory := NewBoolFactory[FeatureFlag]()
//	darkMode := factory.New(true)
//	notifications := factory.New(false)
func NewBoolFactory[T EnumBool[T]]() BoolFactory[T] {
	return BoolFactory[T]{}
}
