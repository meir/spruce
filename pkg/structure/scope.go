package structure

type Scope struct {
	variables map[string]Variable
}

func NewScope() *Scope {
	return &Scope{
		map[string]Variable{},
	}
}

func NewScopeWithParent(parent *Scope) *Scope {
	o := NewScope()
	for k, v := range parent.variables {
		o.variables[k] = v
	}
	return o
}

func (s *Scope) Get(name string) Variable {
	if v, ok := s.variables[name]; ok {
		return v
	}
	return nil
}

func (s *Scope) Set(name string, value Variable) {
	s.variables[name] = value
}
