package spruce

import "reflect"

type Scope struct {
	Parent    *Scope
	Variables map[string]Variable
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent:    parent,
		Variables: make(map[string]Variable),
	}
}

func (s *Scope) Get(name string) Variable {
	if v, ok := s.Variables[name]; ok {
		return v
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s *Scope) GetByType(t Variable) []Variable {
	out := []Variable{}
	for _, v := range s.Variables {
		if reflect.TypeOf(v) == reflect.TypeOf(t) {
			out = append(out, v)
		}
	}
	return out
}

func (s *Scope) Set(name string, v any) {
	if variable, exists := s.Variables[name]; exists {
		variable.Set(v)
	}
	panic("variable not defined")

}

func (s *Scope) Define(name string, v Variable) {
	if _, exists := s.Variables[name]; exists {
		panic("variable already defined")
	}

	s.Variables[name] = v
}
