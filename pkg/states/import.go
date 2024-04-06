package states

import "github.com/meir/spruce/pkg/structure"

type ImportAST struct {
	value structure.AST
}

func (e *ImportAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return true
}

func (e *ImportAST) String(self *structure.ASTWrapper) string {
	return ""
}

type ImportNode struct{}

func NewImportNode() *ImportNode {
	return &ImportNode{}
}

func (e *ImportNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
	}
}

func (e *ImportNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.PeekNext(2)
	if t == nil {
		return 0, nil, nil
	}

	if t.Str == "@import" {
		return structure.STATE_IMPORT, &ImportAST{}, scope
	}
	return 0, nil, nil
}
