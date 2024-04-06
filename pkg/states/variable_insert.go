package states

import "github.com/meir/spruce/pkg/structure"

type VariableInsertAST struct {
	variable string
}

func (e *VariableInsertAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	return true
}

func (e VariableInsertAST) String(self *structure.ASTWrapper) string {
	variable := self.Scope.Get(e.variable)
	if variable == nil {
		return ""
	}
	return variable.String()
}

type VariableInsertNode struct{}

func NewVariableInsertNode() *VariableInsertNode {
	return &VariableInsertNode{}
}

func (e *VariableInsertNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_CONTAINER,
		structure.STATE_ELEMENT_ATTRIBUTE,
	}
}

func (e *VariableInsertNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if t.Equals("$") {
		key := ts.PeekActual(1)
		return structure.STATE_VARIABLE_INSERT, &VariableInsertAST{
			variable: key.Str,
		}, scope
	}
	return 0, nil, nil
}
