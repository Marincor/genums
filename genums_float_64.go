package genums

// EnumFloat64 is a specialized enum interface for float64-based enumerations.
// It extends BaseEnum with float64 as the underlying value type, providing
// type-safe floating-point enum functionality.
//
// Type parameters:
//   - T: The concrete enum type that implements this interface
type EnumFloat64[T any] interface {
	BaseEnum[T, float64]
}

// enumFloat64Impl is the internal implementation of the EnumFloat64 interface.
// It embeds enumImpl with float64 as the value type to provide all
// floating-point-specific enum functionality.
//
// Type parameters:
//   - T: The concrete enum type (used for type safety)
type enumFloat64Impl[T any] struct {
	enumImpl[T, float64]
}

// NewEnumFloat64 creates a new float64-based enum instance with the given value.
// This is the primary constructor for creating floating-point enum values.
//
// Type parameters:
//   - T: The target enum type to create
//
// Parameters:
//   - value: The float64 value to wrap in the enum
//
// Returns:
//   - A new EnumFloat64 instance containing the specified value
//
// Example:
//
//	type Temperature EnumFloat64[Temperature]
//	freezing := NewEnumFloat64[Temperature](0.0)
//	boiling := NewEnumFloat64[Temperature](100.0)
func NewEnumFloat64[T any](value float64) EnumFloat64[T] {
	return enumFloat64Impl[T]{enumImpl[T, float64]{value: value}}
}

// Float64Factory is a factory type that provides a convenient way to create
// float64-based enum instances. It implements a factory pattern for enums,
// allowing for more flexible instantiation.
//
// Type parameters:
//   - T: The enum type to be created (must implement EnumFloat64[T])
type Float64Factory[T EnumFloat64[T]] struct{}

// New creates a new enum instance of type T with the specified float64 value.
// This method allows the factory to produce properly typed enum instances
// using the underlying NewEnumFloat64 constructor.
//
// Parameters:
//   - value: The float64 value for the new enum instance
//
// Returns:
//   - A new enum instance of type T with the given value
//
// Example:
//
//	factory := NewFloat64Factory[Temperature]()
//	freezing := factory.New(0.0)
//	boiling := factory.New(100.0)
func (f Float64Factory[T]) New(value float64) T {
	return any(NewEnumFloat64[T](value)).(T)
}

// NewFloat64Factory creates a new Float64Factory instance for the specified enum type.
// This function initializes a factory that can be used to create multiple enum
// instances of the same type.
//
// Type parameters:
//   - T: The enum type the factory will create (must implement EnumFloat64[T])
//
// Returns:
//   - A new Float64Factory instance for type T
//
// Example:
//
//	type Temperature EnumFloat64[Temperature]
//	factory := NewFloat64Factory[Temperature]()
//	freezing := factory.New(0.0)
//	roomTemp := factory.New(25.0)
//	boiling := factory.New(100.0)
func NewFloat64Factory[T EnumFloat64[T]]() Float64Factory[T] {
	return Float64Factory[T]{}
}
