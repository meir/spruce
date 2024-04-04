package structure

import "strings"

type AST interface {
	Next(t *Tokenizer, self *ASTWrapper) bool
	String(self *ASTWrapper) string
}

type ASTWrapper struct {
	Ast      AST
	Children []*ASTWrapper
	state    State
	Scope    *Scope
}

func (ast *ASTWrapper) JoinChildren() string {
	str := []string{}
	for _, child := range ast.Children {
		str = append(str, child.Ast.String(child))
	}
	return strings.Join(str, "")
}
