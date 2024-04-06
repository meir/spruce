package variables

import "github.com/meir/spruce/pkg/structure"

type ASTsVariable struct {
	Asts []*structure.ASTWrapper
}

func NewASTsVariable(asts []*structure.ASTWrapper) *ASTsVariable {
	return &ASTsVariable{Asts: asts}
}

func (a *ASTsVariable) Get() any {
	return a.Asts
}

func (a *ASTsVariable) Set(value any) {
	if value, ok := value.([]*structure.ASTWrapper); ok {
		a.Asts = value
	}
}

func (a *ASTsVariable) String() string {
	output := ""
	for _, ast := range a.Asts {
		output += ast.Ast.String(ast)
	}
	return output
}
