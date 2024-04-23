package types

import "fmt"

type Int struct {
	val int
}

func (i *Int) Value() any {
	return i.val
}

func (i *Int) Set(v any) {
	if v, ok := v.(int); ok {
		i.val = v
	}
	panic(fmt.Sprintf("expected value type 'int' got value %v", v))
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.val)
}
