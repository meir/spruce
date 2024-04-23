package types

import "fmt"

type Bool struct {
	val bool
}

func (b *Bool) Value() any {
	return b.val
}

func (b *Bool) Set(v any) {
	if v, ok := v.(bool); ok {
		b.val = v
	}
	panic(fmt.Sprintf("expected value type 'bool' got value %v", v))
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.val)
}
