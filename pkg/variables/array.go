package variables

import "fmt"

type ArrayVariable[V any] struct {
	array []V
}

func NewArrayVariable[V any](array []V) *ArrayVariable[V] {
	return &ArrayVariable[V]{array: array}
}

func (a *ArrayVariable[V]) Get() any {
	return a.array
}

func (a *ArrayVariable[V]) Set(value any) {
	if v, ok := value.([]V); ok {
		a.array = v
	}
}

func (a *ArrayVariable[V]) String() string {
	return fmt.Sprint(a.array)
}
