package states

import (
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type VariableAST struct {
	key string
}

func (v *VariableAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	self.Scope.Set(v.key, variables.NewASTsVariable(self.Children))
	return len(self.Children) > 0
}

func (v *VariableAST) String(self *structure.ASTWrapper) string {
	return ""
}

type VariableNode struct{}

func NewVariableNode() *VariableNode {
	return &VariableNode{}
}

func (v *VariableNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
	}
}

func (v *VariableNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.PeekActual(1)
	if t.Equals("=") {
		key := ts.Current()
		ts.SkipTo(t)
		return structure.STATE_VARIABLE_DEFINITION, &VariableAST{
			key: key.Str,
		}, scope
	}
	return 0, nil, nil
}
