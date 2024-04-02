package variables

type StringVariable struct {
	str string
}

func NewStringVariable(str string) *StringVariable {
	return &StringVariable{str: str}
}

func (s *StringVariable) Get() any {
	return s.str
}

func (s *StringVariable) Set(value any) {
	if value, ok := value.(string); ok {
		s.str = value
	}
}

func (s *StringVariable) String() string {
	return s.str
}
