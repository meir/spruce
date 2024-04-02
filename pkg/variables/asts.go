package variables

import "github.com/meir/spruce/pkg/structure"

type ASTsVariable struct {
	asts []*structure.ASTWrapper
}

func NewASTsVariable(asts []*structure.ASTWrapper) *ASTsVariable {
	return &ASTsVariable{asts: asts}
}

func (a *ASTsVariable) Get() any {
	return a.asts
}

func (a *ASTsVariable) Set(value any) {
	if value, ok := value.([]*structure.ASTWrapper); ok {
		a.asts = value
	}
}

func (a *ASTsVariable) String() string {
	return structure.JoinChildren(a.asts)
}
