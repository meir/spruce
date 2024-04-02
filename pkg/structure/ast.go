package structure

type AST interface {
	Next(t *Tokenizer) bool
	String([]*ASTWrapper) string
}

type ASTWrapper struct {
	Ast      AST
	Children []*ASTWrapper
	state    State
	scope    *Scope
}
