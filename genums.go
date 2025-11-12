// Package genums provides a generic enum implementation for Go, allowing type-safe
// enumeration values with comparison, serialization, and value retrieval capabilities.
package genums

import (
	"fmt"
)

// BaseEnum defines the core interface that all enum types must implement.
// It provides methods for equality comparison, value retrieval, and JSON marshaling.
//
// Type parameters:
//   - T: The concrete enum type that implements this interface
//   - V: The underlying value type (must be comparable)
type BaseEnum[T any, V comparable] interface {
	// IsEqual checks if this enum value is equal to another enum value.
	// Returns true if both enums have the same underlying value, false otherwise.
	IsEqual(other T) bool

	// Value returns the underlying value of the enum.
	Value() V

	// MarshalJSON serializes the enum to JSON format.
	// Returns the JSON byte representation and any error that occurred.
	MarshalJSON() ([]byte, error)
}

// enumImpl is the internal implementation of the BaseEnum interface.
// It stores the underlying value and provides all required functionality.
//
// Type parameters:
//   - T: The concrete enum type (not used directly in the implementation)
//   - V: The underlying value type (must be comparable)
type enumImpl[T any, V comparable] struct {
	value V
}

// Value returns the underlying value stored in this enum instance.
func (e enumImpl[T, V]) Value() V {
	return e.value
}

// IsEqual compares this enum with another enum for equality.
// It performs a nil check first, then verifies that the other value
// implements the Value() method and compares the underlying values.
//
// Parameters:
//   - other: The enum value to compare against
//
// Returns:
//   - true if both enums have equal underlying values
//   - false if other is nil, doesn't implement Value(), or has a different value
func (e enumImpl[T, V]) IsEqual(other T) bool {
	// Check if the other value is nil
	if any(other) == nil {
		return false
	}

	// Attempt to extract the Value() method from the other enum
	if otherEnum, ok := any(other).(interface{ Value() V }); ok {
		return e.value == otherEnum.Value()
	}

	return false
}

// String returns a string representation of the enum's underlying value.
// Uses fmt.Sprintf to convert the value to a string.
func (e enumImpl[T, V]) String() string {
	return fmt.Sprintf("%v", e.value)
}

// MarshalJSON implements the json.Marshaler interface for JSON serialization.
// The enum value is converted to a string and wrapped in double quotes
// to produce valid JSON string output.
//
// Returns:
//   - A byte slice containing the JSON-encoded string representation
//   - nil error (this implementation never fails)
func (e enumImpl[T, V]) MarshalJSON() ([]byte, error) {
	stringValue := fmt.Sprintf("%v", e.value)

	return []byte(`"` + stringValue + `"`), nil
}
