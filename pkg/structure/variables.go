package structure

type Variable interface {
	Get() any
	Set(any)
	String() string
}

func Get[V any](v Variable) (V, bool) {
	if v, ok := v.Get().(V); ok {
		return v, true
	}

	var s V
	return s, false
}

func Set(v Variable, value any) {
	v.Set(value)
}
