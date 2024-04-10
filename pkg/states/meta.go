package states

import (
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type MetaAST struct{}

func (m *MetaAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return len(self.Children) == 1
}

func (m *MetaAST) String(self *structure.ASTWrapper) string {
	return ""
}

type MetaNode struct{}

func NewMetaNode() *MetaNode {
	return &MetaNode{}
}

func (m *MetaNode) States() []structure.State {
	return []structure.State{
		structure.STATE_AT_STATEMENT,
	}
}

func (m *MetaNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t == nil {
		return 0, nil, nil
	}

	meta_scope := structure.NewScope()
	meta_scope.Set("attributes", variables.NewMapVariable(map[string]any{}))

	if t.Str == "meta" {
		return structure.STATE_META, &MetaAST{}, meta_scope
	}
	return 0, nil, nil
}
