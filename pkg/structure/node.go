package structure

type Node interface {
	Active(l *Tokenizer, scope *Scope) (State, AST, *Scope)
	States() []State
}
