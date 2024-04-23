package spruce

import (
	"context"
)

type Lexer struct {
	file *File

	history []context.Context
}

func NewLexer(file *File) (*Lexer, error) {
	return &Lexer{
		file:    file,
		history: []context.Context{},
	}, nil
}

func (l *Lexer) Parse(ctx context.Context) context.Context {
	tokenizer := GetTokenizer(ctx)
	l.Up(ctx)

MainLoop:
	for tokenizer.HasNext() {
		err := tokenizer.Next()
		if err != nil {
			panic(err)
		}

		nodes := GetNodesByState(GetState(l.Current()))

		for _, node := range nodes {
			nctx, active := node.Activate(l.Current())
			if active {
				nctx = SetChildren(nctx, &[]context.Context{})
				l.Up(nctx)
				continue MainLoop
			}
		}

		l.ContinueAST()
	}

	return l.Current()
}

func (l *Lexer) ContinueAST() {
	ctx := l.Current()
	if GetAST(ctx) != nil {
		done, err := GetAST(l.Current()).Build(l.Current())
		if err != nil {
			panic(err)
		}

		if done {
			l.Down()
			*GetChildren(l.Current()) = append(*GetChildren(l.Current()), ctx)
		}
	}
}

func (l *Lexer) Current() context.Context {
	return l.history[len(l.history)-1]
}

func (l *Lexer) Up(ctx context.Context) {
	l.history = append(l.history, ctx)
}

func (l *Lexer) Down() {
	l.history = l.history[:len(l.history)-1]
}
