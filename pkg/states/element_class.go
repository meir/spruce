package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ElementClassAST struct {
	class string
}

func (e *ElementClassAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	class := self.Scope.Get("class")
	if m, ok := structure.Get[string](class); ok {
		if m == "" {
			m = e.class
		} else {
			m += " " + e.class
		}
		structure.Set(class, m)
	}
	return true
}

func (e ElementClassAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ElementClassNode struct{}

func NewElementClassNode() *ElementClassNode {
	return &ElementClassNode{}
}

func (e *ElementClassNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
	}
}

func (e *ElementClassNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.PeekNext(2)
	if t.EqualsRegexp(`\..+`) {
		return structure.STATE_ELEMENT_CLASS, &ElementClassAST{
			class: t.Str[1:],
		}
	}
	return 0, nil
}
