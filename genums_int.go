package genums

// EnumInt is a specialized enum interface for integer-based enumerations.
// It extends BaseEnum with int as the underlying value type, providing
// type-safe integer enum functionality.
//
// Type parameters:
//   - T: The concrete enum type that implements this interface
type EnumInt[T any] interface {
	BaseEnum[T, int]
}

// enumIntImpl is the internal implementation of the EnumInt interface.
// It embeds enumImpl with int as the value type to provide all
// integer-specific enum functionality.
//
// Type parameters:
//   - T: The concrete enum type (used for type safety)
type enumIntImpl[T any] struct {
	enumImpl[T, int]
}

// NewEnumInt creates a new integer-based enum instance with the given value.
// This is the primary constructor for creating integer enum values.
//
// Type parameters:
//   - T: The target enum type to create
//
// Parameters:
//   - value: The integer value to wrap in the enum
//
// Returns:
//   - A new EnumInt instance containing the specified value
//
// Example:
//
//	type Priority EnumInt[Priority]
//	high := NewEnumInt[Priority](1)
//	medium := NewEnumInt[Priority](2)
//	low := NewEnumInt[Priority](3)
func NewEnumInt[T any](value int) EnumInt[T] {
	return enumIntImpl[T]{enumImpl[T, int]{value: value}}
}
