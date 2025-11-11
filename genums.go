package genums

type baseEnum[T any] interface {
	IsEqual(provider T) bool
	// MarshalJSON returns the JSON representation of the enum.
	MarshalJSON() ([]byte, error)
	Add(value string) baseEnum[T]
}

type EnumString[T baseEnum[T]] interface {
	baseEnum[T]
}

func NewFactory[T any]() baseEnum[T] {
	var newInstance EnumStringImpl[T]

	return newInstance

}

type EnumStringImpl[T any] string

func (e EnumStringImpl[T]) MarshalJSON() ([]byte, error) {
	return []byte("\"" + string(e) + "\""), nil
}

func (e EnumStringImpl[T]) Raw() string {
	return string(e)
}

func (e EnumStringImpl[T]) IsEqual(rawValue T) bool {
	if any(rawValue) == nil {
		return false
	}

	return e.Raw() == any(rawValue)
}

func (e EnumStringImpl[T]) Add(value string) baseEnum[T] {
	return EnumStringImpl[T](value)
}
