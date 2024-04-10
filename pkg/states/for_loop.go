package states

import "github.com/meir/spruce/pkg/structure"

type ForLoopAST struct {
	variable string
}

func (f *ForLoopAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return false
}

func (f *ForLoopAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ForLoopNode struct{}

func NewForLoopNode() *ForLoopNode {
	return &ForLoopNode{}
}

func (f *ForLoopNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_CONTAINER,
	}
}

func (f *ForLoopNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t == nil {
		return 0, nil, nil
	}
	if t.Str == "for" {
		return structure.STATE_FOR_LOOP, &ForLoopAST{}, scope
	}
	return 0, nil, nil
}
