package structure

import (
	"fmt"
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
			if state, ast, scope := node.Active(l.tokens, scope); ast != nil {
				l.state = state
				asts = append(asts, &ASTWrapper{
					Ast:      ast,
					Children: []*ASTWrapper{},
					state:    state,
					Scope:    scope,
				})
				continue TokenLoop
			}
		}

		if len(asts) == 0 {
			curr := l.tokens.Current()
			if curr.IsEmpty() {
				continue TokenLoop
			}
			panic(fmt.Errorf("unexpected token: %s at %d:%d (state: %d)", curr.Str, curr.Line+1, curr.Start+1, l.state))
		}

		currentAst := asts[len(asts)-1]
		scope = currentAst.Scope
		endAst := currentAst.Ast.Next(l.tokens, currentAst)

		if endAst {
			asts = asts[:len(asts)-1]

			if len(asts) > 0 {
				parent := asts[len(asts)-1]
				parent.Children = append(parent.Children, currentAst)
				l.state = parent.state
			} else {
				root = append(root, currentAst)
				l.state = STATE_ROOT
			}
		}
	}

	return root
}

func (l *Lexer) Format(ast []*ASTWrapper) string {
	builder := strings.Builder{}
	for _, a := range ast {
		builder.WriteString(a.Ast.String(a))
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
