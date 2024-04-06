package states

import (
	"github.com/meir/spruce/pkg/structure"
	"github.com/meir/spruce/pkg/variables"
)

type TemplateAST struct {
	template *variables.ASTsVariable
}

func (e *TemplateAST) Next(ts *structure.Tokenizer, self *structure.ASTWrapper) bool {
	hasContainer := false
	if len(self.Children) == 0 {
		return false
	}

	if _, ok := self.Children[len(self.Children)-1].Ast.(*ContainerAST); ok {
		hasContainer = true
		asts := []*structure.ASTWrapper{}
		for _, ast := range self.Children[len(self.Children)-1].Children {
			switch ast.Ast.(type) {
			// TODO: Ignore template definition types once implemented
			default:
				asts = append(asts, ast)
			}
		}

		self.Scope.Set("@", variables.NewASTsVariable(asts))
	}

	return hasContainer
}

func (e TemplateAST) String(self *structure.ASTWrapper) string {
	self.Scope.AppendOnChildren(e.template.Asts)
	return e.template.String()
}

type TemplateNode struct{}

func NewTemplateNode() *TemplateNode {
	return &TemplateNode{}
}

func (e *TemplateNode) States() []structure.State {
	return []structure.State{
		structure.STATE_ROOT,
		structure.STATE_CONTAINER,
	}
}

func (e *TemplateNode) Active(ts *structure.Tokenizer, scope *structure.Scope) (structure.State, structure.AST, *structure.Scope) {
	t := ts.Current()
	if variable := scope.Get(t.Str); variable != nil {
		if astVariable, ok := variable.(*variables.ASTsVariable); ok {
			return structure.STATE_TEMPLATE, &TemplateAST{
				template: astVariable,
			}, structure.NewScopeWithParent(scope)
		}
	}
	return 0, nil, nil
}
