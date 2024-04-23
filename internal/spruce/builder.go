package spruce

import (
	"context"
)

func Parse(file string) (context.Context, error) {
	ctx := context.Background()

	scope := NewScope(nil)
	f, err := NewFile(file)
	if err != nil {
		return nil, err
	}

	tokenizer, err := NewTokenizer(f)
	if err != nil {
		return nil, err
	}

	lexer, err := NewLexer(f)
	if err != nil {
		return nil, err
	}

	state := STATE_ROOT
	children := &[]context.Context{}

	ctx = SetScope(ctx, scope)
	ctx = SetFile(ctx, f)
	ctx = SetTokenizer(ctx, tokenizer)
	ctx = SetLexer(ctx, lexer)
	ctx = SetState(ctx, state)
	ctx = SetChildren(ctx, children)

	return lexer.Parse(ctx), nil
}

func String(ctx context.Context) string {
	return BuildChildren(ctx)
}
