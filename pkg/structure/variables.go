package structure

type Scope map[string]Variable

func NewScope() *Scope {
	return &Scope{}
}

func NewScopeWithParent(parent *Scope) *Scope {
	o := NewScope()
	for k, v := range *parent {
		(*o)[k] = v
	}
	return o
}

func (s *Scope) Get(name string) Variable {
	if v, ok := (*s)[name]; ok {
		return v
	}
	return nil
}

func (s *Scope) Set(name string, value Variable) {
	(*s)[name] = value
}

type Variable interface {
	String() string
}
