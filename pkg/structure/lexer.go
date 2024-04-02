package structure

import (
	"slices"
	"strings"
)

type Lexer struct {
	tokens *Tokenizer

	state State
	nodes []Node
}

func NewLexer(tokenizer *Tokenizer, nodes []Node) *Lexer {
	return &Lexer{
		tokens: tokenizer,
		nodes:  nodes,
	}
}

func (l *Lexer) Parse() []*ASTWrapper {
	asts := []*ASTWrapper{}
	root := []*ASTWrapper{}

	scope := NewScope()

TokenLoop:
	for l.tokens.Next() {
		nodes := l.stateNodes()

		for _, node := range nodes {
			if state, ast := node.Active(l.tokens, scope); ast != nil {
				l.state = state
				asts = append(asts, &ASTWrapper{
					Ast:      ast,
					Children: []*ASTWrapper{},
					state:    state,
					scope:    NewScopeWithParent(scope),
				})
				continue TokenLoop
			}
		}

		if len(asts) != 0 {
			current := asts[len(asts)-1]
			scope = current.scope
			if end := current.Ast.Next(l.tokens); end {
				if len(asts) > 1 {
					parent := asts[len(asts)-2]
					parent.Children = append(parent.Children, current)
					l.state = parent.state
					asts = asts[:len(asts)-1]
					continue TokenLoop
				}
				asts = asts[:len(asts)-1]
				if len(asts) == 0 {
					root = append(root, current)
					l.state = STATE_ROOT
				}
			}
		}
	}

	return root
}

func (l *Lexer) Format(ast []*ASTWrapper) string {
	builder := strings.Builder{}
	for _, a := range ast {
		builder.WriteString(a.Ast.String(a.Children))
	}
	return builder.String()
}

func (l *Lexer) stateNodes() []Node {
	n := []Node{}
	for _, node := range l.nodes {
		if slices.Contains(node.States(), l.state) {
			n = append(n, node)
		}
	}
	return n
}
