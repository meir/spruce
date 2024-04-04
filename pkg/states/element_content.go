package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ElementContentAST struct{}

func (e *ElementContentAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	t := ts.Current()
	switch t.Str {
	case "}":
		ts.Skip(-1)
		return true
	default:
		return false
	}
}

func (e ElementContentAST) String(self *structure.ASTWrapper) string {
	return self.JoinChildren()
}

type ElementContentNode struct{}

func NewElementContentNode() *ElementContentNode {
	return &ElementContentNode{}
}

func (e *ElementContentNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
	}
}

func (e *ElementContentNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST) {
	t := ts.Current()
	switch t.Str {
	case "{":
		return structure.STATE_ELEMENT_CONTENT, &ElementContentAST{}
	default:
		return 0, nil
	}
}
