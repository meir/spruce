package states

import "github.com/meir/spruce/pkg/structure"

type AtStatementAST struct{}

func (e *AtStatementAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return len(self.Children) > 0
}

func (e *AtStatementAST) String(self *structure.ASTWrapper) string {
	return ""
}

type AtStatementNode struct{}

func NewAtStatementNode() *AtStatementNode {
	return &AtStatementNode{}
}

func (e *AtStatementNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
	}
}

func (e *AtStatementNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t == nil {
		return 0, nil, nil
	}
	if t.Str == "@" {
		return structure.STATE_AT_STATEMENT, &AtStatementAST{}, scope
	}
	return 0, nil, nil
}
