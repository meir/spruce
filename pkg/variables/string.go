package variables

import "fmt"

type StringVariable struct {
	str any
}

func NewStringVariable(str any) *StringVariable {
	return &StringVariable{str: str}
}

func (s *StringVariable) Get() any {
	return s.str
}

func (s *StringVariable) Set(value any) {
	s.str = value
}

func (s *StringVariable) String() string {
	return fmt.Sprint(s.str)
}
