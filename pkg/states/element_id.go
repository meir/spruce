package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ElementIdAST struct {
	id string
}

func (e *ElementIdAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	id := self.Scope.Get("id")
	id.Set(e.id)
	return true
}

func (e ElementIdAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ElementIdNode struct{}

func NewElementIdNode() *ElementIdNode {
	return &ElementIdNode{}
}

func (e *ElementIdNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
	}
}

func (e *ElementIdNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.PeekNext(2)
	if t.EqualsRegexp(`#.+`) {
		ts.Skip(1)
		return structure.STATE_ELEMENT_ID, &ElementIdAST{
			id: t.Str[1:],
		}, scope
	}
	return 0, nil, nil
}
