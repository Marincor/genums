# Genums

A type-safe, generic enum library for Go that provides a flexible and elegant way to create enumerations with various underlying types.

## Features

- **Type-Safe**: Leverages Go generics for compile-time type safety
- **Multiple Types**: Support for `string`, `int`, `float64`, and `bool` enums
- **JSON Serialization**: Built-in JSON marshaling support
- **Comparison**: Safe equality comparison between enum values
- **Factory Pattern**: Optional factory pattern for convenient enum creation
- **Zero Dependencies**: Pure Go implementation with no external dependencies

## Installation

```bash
go get github.com/marincor/genums
```

## Quick Start

### String Enums

```go
package main

import "github.com/marincor/genums"

type Status genums.EnumString[Status]

var (
    StatusActive   = genums.NewEnumString[Status]("active")
    StatusInactive = genums.NewEnumString[Status]("inactive")
    StatusPending  = genums.NewEnumString[Status]("pending")
)

func main() {
    currentStatus := StatusActive
    
    // Get the underlying value
    fmt.Println(currentStatus.Value()) // Output: active
    
    // Compare enums
    if currentStatus.IsEqual(StatusActive) {
        fmt.Println("Status is active")
    }
    
    // JSON marshaling
    jsonData, _ := json.Marshal(currentStatus)
    fmt.Println(string(jsonData)) // Output: "active"
}
```

### Integer Enums

```go
type Priority genums.EnumInt[Priority]

var (
    PriorityLow    = genums.NewEnumInt[Priority](1)
    PriorityMedium = genums.NewEnumInt[Priority](2)
    PriorityHigh   = genums.NewEnumInt[Priority](3)
)

func main() {
    task := PriorityHigh
    fmt.Println(task.Value()) // Output: 3
}
```

### Float64 Enums

```go
type Temperature genums.EnumFloat64[Temperature]

var (
    Freezing = genums.NewEnumFloat64[Temperature](0.0)
    RoomTemp = genums.NewEnumFloat64[Temperature](25.0)
    Boiling  = genums.NewEnumFloat64[Temperature](100.0)
)

func main() {
    current := RoomTemp
    fmt.Println(current.Value()) // Output: 25.0
}
```

### Boolean Enums

```go
type FeatureFlag genums.EnumBool[FeatureFlag]

var (
    Enabled  = genums.NewEnumBool[FeatureFlag](true)
    Disabled = genums.NewEnumBool[FeatureFlag](false)
)

func main() {
    darkMode := Enabled
    fmt.Println(darkMode.Value()) // Output: true
}
```

## Using Factories

Factories provide an alternative way to create enum instances, useful for dynamic enum creation:

### String Factory

```go
type Status genums.EnumString[Status]

func main() {
    factory := genums.NewStringFactory[Status]()
    
    active := factory.New("active")
    inactive := factory.New("inactive")
    
    fmt.Println(active.Value())   // Output: active
    fmt.Println(inactive.Value()) // Output: inactive
}
```

### Float64 Factory

```go
type Temperature genums.EnumFloat64[Temperature]

func main() {
    factory := genums.NewFloat64Factory[Temperature]()
    
    freezing := factory.New(0.0)
    boiling := factory.New(100.0)
    
    fmt.Println(freezing.Value()) // Output: 0.0
}
```

### Bool Factory

```go
type FeatureFlag genums.EnumBool[FeatureFlag]

func main() {
    factory := genums.NewBoolFactory[FeatureFlag]()
    
    enabled := factory.New(true)
    disabled := factory.New(false)
}
```

## Core Interface

All enum types implement the `BaseEnum` interface:

```go
type BaseEnum[T any, V comparable] interface {
    // Check equality with another enum
    IsEqual(other T) bool
    
    // Get the underlying value
    Value() V
    
    // JSON marshaling support
    MarshalJSON() ([]byte, error)
}
```

## Supported Types

| Type | Constructor | Factory | Use Case |
|------|------------|---------|----------|
| `EnumString` | `NewEnumString[T](string)` | `NewStringFactory[T]()` | Status codes, categories, labels |
| `EnumInt` | `NewEnumInt[T](int)` | N/A | Priority levels, error codes, IDs |
| `EnumFloat64` | `NewEnumFloat64[T](float64)` | `NewFloat64Factory[T]()` | Measurements, thresholds, percentages |
| `EnumBool` | `NewEnumBool[T](bool)` | `NewBoolFactory[T]()` | Feature flags, toggles, binary states |

## Advanced Usage

### JSON Serialization

```go
type Status genums.EnumString[Status]

var StatusActive = genums.NewEnumString[Status]("active")

type Response struct {
    Status Status `json:"status"`
}

func main() {
    resp := Response{Status: StatusActive}
    jsonData, _ := json.Marshal(resp)
    fmt.Println(string(jsonData)) // Output: {"status":"active"}
}
```

### Comparison

```go
status1 := genums.NewEnumString[Status]("active")
status2 := genums.NewEnumString[Status]("active")
status3 := genums.NewEnumString[Status]("inactive")

fmt.Println(status1.IsEqual(status2)) // Output: true
fmt.Println(status1.IsEqual(status3)) // Output: false
```

### Type Safety at Compile Time

Genums leverages Go's type system to prevent type mismatches at compile time:

```go
type Status genums.EnumString[Status]
type Role genums.EnumString[Role]

var StatusActive = genums.NewEnumString[Status]("active")
var RoleAdmin = genums.NewEnumString[Role]("admin")

// This will NOT compile - compiler error!
// StatusActive.IsEqual(RoleAdmin) // ❌ Error: cannot use RoleAdmin (type Role) as type Status

// This compiles correctly
StatusActive.IsEqual(StatusActive) // ✅ OK
```

The compiler ensures you can't accidentally compare or mix enums of different types, catching bugs before runtime.

### String Representation

```go
priority := genums.NewEnumInt[Priority](1)
fmt.Println(priority.String()) // Output: 1
```

## Design Patterns

### Predefined Constants

```go
type Role genums.EnumString[Role]

var (
    RoleAdmin     = genums.NewEnumString[Role]("admin")
    RoleUser      = genums.NewEnumString[Role]("user")
    RoleModerator = genums.NewEnumString[Role]("moderator")
)
```

### Dynamic Creation

```go
func getRoleFromString(s string) Role {
    factory := genums.NewStringFactory[Role]()
    return factory.New(s)
}
```

### Function Parameters

```go
func checkPermission(role Role) bool {
    return role.IsEqual(RoleAdmin) || role.IsEqual(RoleModerator)
}
```

## Why Genums?

- **Compile-Time Type Safety**: The library enforces type constraints at compile time, making it impossible to mix different enum types or pass incorrect types. The Go compiler will catch these errors before your program runs, ensuring type correctness from the start.
- **Clean API**: Simple and intuitive interface
- **Standard Compliance**: Works seamlessly with Go's `json` package
- **Flexible**: Support for multiple underlying types
- **No Reflection**: Performance-friendly implementation

## Requirements

- Go 1.18 or higher (for generics support)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Author

Marincor

## Changelog

### v0.0.1
- Initial release
- Support for String, Int, Float64, and Bool enums
- Factory pattern implementation
- JSON marshaling support
- Equality comparison