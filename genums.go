package genums

type enum struct {
	Name  string
	Value []string
}

func New(name string, values ...string) *enum {
	return &enum{
		Name:  name,
		Value: values,
	}
}
