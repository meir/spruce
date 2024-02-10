package structure

import "slices"

type Lexer struct {
	tokens *Tokenizer

	state State
	nodes []Node
}

type State int64

type AST interface {
	Type() string
	String() string
}

type Node struct {
	Active func(l *Token) bool
	Next   func(l *Token) (*State, *AST)

	States []State
}

func NewLexer(tokenizer *Tokenizer) *Lexer {
	return &Lexer{
		tokens: tokenizer,
	}
}

func (l *Lexer) Lex() {
}

func (l *Lexer) stateNodes() []Node {
	n := []Node{}
	for _, node := range l.nodes {
		if slices.Contains(node.States, l.state) {
			n = append(n, node)
		}
	}
	return n
}
