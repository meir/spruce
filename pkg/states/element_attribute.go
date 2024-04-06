package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ElementAttributeAST struct {
	key string
}

func (e *ElementAttributeAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	attributes := self.Scope.Get("attributes")
	if m, ok := structure.Get[map[string]any](attributes); ok {
		m[e.key] = self.JoinChildren()
		structure.Set(attributes, m)
	}
	return len(self.Children) > 0
}

func (e *ElementAttributeAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ElementAttributeNode struct{}

func NewElementAttributeNode() *ElementAttributeNode {
	return &ElementAttributeNode{}
}

func (e *ElementAttributeNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_CONTAINER,
	}
}

func (e *ElementAttributeNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.PeekActual(1)
	if t.Equals("=") {
		key := ts.Current().Str
		return structure.STATE_ELEMENT_ATTRIBUTE, &ElementAttributeAST{
			key: key,
		}, scope
	}
	return 0, nil, nil
}
