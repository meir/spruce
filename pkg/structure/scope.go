package structure

import "fmt"

type Scope struct {
	variables map[string]Variable
}

func NewScope() *Scope {
	return &Scope{
		map[string]Variable{},
	}
}

func (s *Scope) Dump() {
	for k, v := range s.variables {
		fmt.Printf("%s: %s\n", k, v.String())
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

func (s *Scope) Merge(s2 *Scope) {
	for k, v := range s2.variables {
		s.variables[k] = v
	}
}

func (s *Scope) AppendOnChildren(children []*ASTWrapper) {
	for _, child := range children {
		child.Scope.Merge(s)
		child.Scope.AppendOnChildren(child.Children)
	}
}
