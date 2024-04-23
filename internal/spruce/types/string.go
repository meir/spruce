package types

import "fmt"

type String struct {
	val string
}

func (s *String) Value() any {
	return s.val
}

func (s *String) Set(v any) {
	if v, ok := v.(string); ok {
		s.val = v
	}
	panic(fmt.Sprintf("expected value type 'str' got value %v", v))
}

func (s *String) String() string {
	return s.val
}
