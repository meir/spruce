package types

import "fmt"

type Attr struct {
	key any
	val any
}

func NewAttr(key, value any) *Attr {
	return &Attr{key, value}
}

func (b *Attr) Value() any {
	return []any{b.key, b.val}
}

func (b *Attr) Set(v any) {
	if v, ok := v.([]any); ok {
		if len(v) == 2 {
			b.key = v[0]
			b.val = v[1]
			return
		}
	}
	panic(fmt.Sprintf("expected value type '[string, string]' got value %v", v))
}

func (b *Attr) String() string {
	key := fmt.Sprintf("%v", b.key)
	val := fmt.Sprintf("%v", b.val)
	return fmt.Sprintf("%s=\"%s\"", key, val)
}
