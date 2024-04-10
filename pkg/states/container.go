package states

import (
	"github.com/meir/spruce/pkg/structure"
)

type ContainerAST struct{}

func (e *ContainerAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	t := ts.Current()
	switch t.Str {
	case "}":
		return true
	default:
		return false
	}
}

func (e ContainerAST) String(self *structure.ASTWrapper) string {
	return self.JoinChildren()
}

type ContainerNode struct{}

func NewContainerNode() *ContainerNode {
	return &ContainerNode{}
}

func (e *ContainerNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ELEMENT,
		structure.STATE_VARIABLE_DEFINITION,
		structure.STATE_TEMPLATE,
		structure.STATE_META,
	}
}

func (e *ContainerNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	switch t.Str {
	case "{":
		return structure.STATE_CONTAINER, &ContainerAST{}, scope
	default:
		return 0, nil, nil
	}
}
